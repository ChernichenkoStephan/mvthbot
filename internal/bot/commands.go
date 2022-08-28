package bot

import (
	"context"
	"fmt"
	"strings"
	"syscall"
	"time"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"

	tele "gopkg.in/telebot.v3"
)

func NewTeleHandler(ctx context.Context, f HandleFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := f(ctx, c)
		return err
	}
}

func (b *Bot) HandleDefault(ctx context.Context, c tele.Context) error {
	dest := c.Sender()

	// To work with empty (example: 2+2) statements only in bot
	if c.Chat().Type == "private" {

		sts, ok := c.Get("statements").([]solving.Statement)
		if !ok {
			return fmt.Errorf("wrong type from middleware %v", sts)
		}

		resp, procErr := b.process(ctx, dest.ID, sts)
		if procErr != nil {
			resp = fmt.Sprintf("Wrong input.\n%s", procErr.Error())
		}

		sendError := c.Send(resp)
		if sendError != nil || procErr != nil {
			return errors.Combine(procErr, sendError)
		}
	}

	return nil
}

// s command
func (b *Bot) HandleSolve(ctx context.Context, c tele.Context) error {

	sts, ok := c.Get("statements").([]solving.Statement)
	if !ok {
		return fmt.Errorf("wrong type from middleware %v", sts)
	}

	if len(sts) == 0 {
		return c.Send(`Give me a statement (example: "/s 2+2" or "/s a=2+2" or just "/s a=4")`)
	}

	resp, procErr := b.process(ctx, c.Sender().ID, sts)
	if procErr != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", procErr.Error())
	}

	sendError := c.Send(resp)
	if sendError != nil || procErr != nil {
		return errors.Combine(procErr, sendError)
	}

	return nil

}

// get comamnd
func (b *Bot) HandleGetVariables(ctx context.Context, c tele.Context) error {
	var resp string

	fmt.Printf("Got args: %v", c.Args())

	if len(c.Args()) == 0 {
		return c.Send(`Give me variable names separated by spases (example: "/g a b c")`)
	}

	vs, servError := b.db.GetVariablesWithNames(ctx, c.Sender().ID, c.Args())
	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		if len(vs) > 0 {
			builder := NewOutputBuilder()
			for n, v := range vs {
				builder.WriteVariable(n)
				builder.WriteValue(v)
				builder.LineBreak()
			}
			resp = builder.String()
		} else {
			resp = `The list is empty, it's time to make some!`
		}
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil

}

// getall comamnd
func (b *Bot) HandleGetAllVariables(ctx context.Context, c tele.Context) error {
	var resp string

	vs, servError := b.db.GetAllVariables(ctx, c.Sender().ID)
	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		if len(vs) > 0 {
			builder := NewOutputBuilder()
			for n, v := range vs {
				builder.WriteVariable(n)
				builder.WriteValue(v)
				builder.LineBreak()
			}
			resp = builder.String()
		} else {
			resp = `The list is empty, it's time to make some!`
		}
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil

}

// del command
func (b *Bot) HandleDeleteVariables(ctx context.Context, c tele.Context) error {
	var resp string

	servError := b.db.DeleteVariablesWithNames(ctx, c.Sender().ID, c.Args())
	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		resp = "Success"
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil

}

// delall comand
func (b *Bot) HandleDeleteAllVariables(ctx context.Context, c tele.Context) error {
	var resp string

	servError := b.db.DeleteAllVariables(ctx, c.Sender().ID)

	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		resp = "Success"
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil
}

// hist command
func (b *Bot) HandleGetHistory(ctx context.Context, c tele.Context) error {
	var resp string

	sts, servError := b.db.GetHistory(ctx, c.Sender().ID)
	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		builder := NewOutputBuilder()
		for _, s := range *sts {
			builder.WriteFull(&s)
		}
		resp = builder.String()
		if resp == "" {
			resp = "empty"
		}
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil
}

// clear command
func (b *Bot) HandleClearAll(ctx context.Context, c tele.Context) error {
	var resp string

	servError := b.db.Clear(ctx, c.Sender().ID)
	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		resp = "Success"
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil
}

// password command
func (b *Bot) HandleGetPassword(ctx context.Context, c tele.Context) error {
	var resp string
	var servError error

	if c.Chat().Type == "private" {

		u, servError := b.db.Get(ctx, c.Sender().ID)
		if servError != nil {
			resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
		} else if u.Password == "" {
			resp = "No password. Generete one with /genpassword command"
		} else {
			resp = u.Password
		}
	} else {
		resp = "Command forbiden, use only in private bot chat"
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil
}

// genpassword command
func (b *Bot) HandleGeneratePassword(ctx context.Context, c tele.Context) error {
	var resp string
	var servError error

	if c.Chat().Type == "private" {

		u, servError := b.db.Get(ctx, c.Sender().ID)
		if servError != nil {
			resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
		} else {
			u.Password = utils.GenPassword(b.conf.PasswordLength)
			resp = fmt.Sprintf("%v", u.Password)
		}

	} else {
		resp = "Command forbiden, use only in private bot chat"
	}

	sendError := c.Send(resp)
	if sendError != nil || servError != nil {
		return errors.Combine(servError, sendError)
	}

	return nil
}

// start command
func (b *Bot) HandleGreatings(ctx context.Context, c tele.Context) error {
	err := c.Send(b.GetCommandText(`start`))
	if err != nil {
		return errors.Wrap(err, "Reply failed")
	}
	return nil
}

// help command
func (b *Bot) HandleHelp(ctx context.Context, c tele.Context) error {

	sendError := c.Send(b.GetCommandText(`help`))
	if sendError != nil {
		return errors.Wrap(sendError, "Reply failed.")
	}

	return nil
}

// For not there is only 2 types, so i can skip value check
func isRoot(c tele.Context) bool {
	return c.Get("RIGHTS") != nil
}

func setTextCommand(c tele.Context, f func(text string)) error {
	resp := `U shuld be admin to make this. (or got wrong key)`

	if isRoot(c) {

		start := strings.Index(c.Message().Text, `#`)
		if start != -1 {
			newHelpText := c.Message().Text[start:]
			f(newHelpText)
			resp = fmt.Sprintf("Success with: %s", newHelpText)
		} else {
			resp = `Start text with "#" char`
		}

	}

	sendError := c.Send(resp)
	if sendError != nil {
		return errors.Wrap(sendError, "Reply failed.")
	}

	return nil
}

// set help text command
func (b *Bot) HandleSetHelp(ctx context.Context, c tele.Context) error {
	return setTextCommand(c, func(text string) {
		b.SetCommandText(`help`, text)
	})
}

// set greetings text command
func (b *Bot) HandleSetGreet(ctx context.Context, c tele.Context) error {
	return setTextCommand(c, func(text string) {
		b.SetCommandText(`start`, text)
	})
}

// set greetings text command
func (b *Bot) HandleAbort(ctx context.Context, c tele.Context) error {
	if isRoot(c) {

		err := c.Send(`Shuting down...`)
		if err != nil {
			return errors.Wrap(err, "Reply failed.")
		}

		syscall.Kill(syscall.Getpid(), syscall.SIGINT)

		time.Sleep(time.Minute)

		err = c.Send(`Shut down not succeed`)
		if err != nil {
			return errors.Wrap(err, "Reply failed.")
		}

	} else {
		err := c.Send(`U shuld be admin to make this. (or got wrong key)`)
		if err != nil {
			return errors.Wrap(err, "Reply failed.")
		}
	}

	return nil

}

func (b Bot) BaseCommands() *[]Command {
	res := &[]Command{
		{
			Meta: tele.Command{
				Text:        "/s",
				Description: "Solves expressions, usage: /s 2+2",
			},
			Handler: b.HandleSolve,
		},
		{
			Meta: tele.Command{
				Text:        "/g",
				Description: "Returns variable(s), usage: /g a",
			},
			Handler: b.HandleGetVariables,
		},
		{
			Meta: tele.Command{
				Text:        "/getall",
				Description: "Returns all variables",
			},
			Handler: b.HandleGetAllVariables,
		},
		{
			Meta: tele.Command{
				Text:        "/d",
				Description: "Deletes variable, usage: /d a",
			},
			Handler: b.HandleDeleteVariables,
		},
		{
			Meta: tele.Command{
				Text:        "/delall",
				Description: "Delete all variables",
			},
			Handler: b.HandleDeleteAllVariables,
		},
		{
			Meta: tele.Command{
				Text:        "/hist",
				Description: "Returns solving history",
			},
			Handler: b.HandleGetHistory,
		},
		{
			Meta: tele.Command{
				Text:        "/clear",
				Description: "Clears expressions history and variables",
			},
			Handler: b.HandleClearAll,
		},
		{
			Meta: tele.Command{
				Text:        "/password",
				Description: "Returns current password for REST API",
			},
			Handler: b.HandleGetPassword,
		},
		{
			Meta: tele.Command{
				Text:        "/genpassword",
				Description: "Generates new password for REST API",
			},
			Handler: b.HandleGeneratePassword,
		},
		{
			Meta: tele.Command{
				Text:        "/help",
				Description: "Outputs detailed description",
			},
			Handler: b.HandleHelp,
		},
		{
			Meta: tele.Command{
				Text:        "/start",
				Description: "Outputs welcome text",
			},
			Handler: b.HandleGreatings,
		},
		{
			Meta: tele.Command{
				Text:        "/seth",
				Description: "Sets help text",
			},
			IsPrivate: true,
			Handler:   b.HandleSetHelp,
		},
		{
			Meta: tele.Command{
				Text:        "/setg",
				Description: "Sets start text",
			},
			IsPrivate: true,
			Handler:   b.HandleSetGreet,
		},
		{
			Meta: tele.Command{
				Text:        "/abort",
				Description: "Stops the service",
			},
			IsPrivate: true,
			Handler:   b.HandleAbort,
		},
	}
	return res
}
