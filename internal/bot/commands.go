package bot

import (
	"context"
	"fmt"

	"emperror.dev/errors"
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

		resp, procErr := b.process(ctx, dest.ID, c.Get("statements"))
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

// start command
func (b *Bot) HandleGreatings(ctx context.Context, c tele.Context) error {
	err := c.Send("Hi, let's go!")
	if err != nil {
		return errors.Wrap(err, "Reply failed")
	}
	return nil
}

// s command
func (b *Bot) HandleSolve(ctx context.Context, c tele.Context) error {

	resp, procErr := b.process(ctx, c.Sender().ID, c.Get("statements"))
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

	vs, servError := b.db.GetVariablesWithNames(ctx, c.Sender().ID, c.Args())
	if servError != nil {
		resp = fmt.Sprintf("Wrong input.\n%s", servError.Error())
	} else {
		builder := NewOutputBuilder()
		for n, v := range vs {
			builder.WriteVariable(n)
			builder.WriteValue(v)
			builder.LineBreak()
		}
		resp = builder.String()
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
		builder := NewOutputBuilder()
		for n, v := range vs {
			builder.WriteVariable(n)
			builder.WriteValue(v)
			builder.LineBreak()
		}
		resp = builder.String()
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
			// TODO fix to config
			u.Password = utils.GenPassword(8)
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

// help command
func (b *Bot) HandleHelp(ctx context.Context, c tele.Context) error {

	resp := "HandleHelp command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())

	sendError := c.Send(resp)
	if sendError != nil {
		return errors.Wrap(sendError, "Reply failed.")
	}

	return nil
}

func (b Bot) BaseCommands() *[]Command {
	res := &[]Command{
		{
			Meta: tele.Command{
				Text:        "/s",
				Description: "Solve command",
			},
			Handler:         b.HandleSolve,
			IsParameterized: true,
		},
		{
			Meta: tele.Command{
				Text:        "/get",
				Description: "Get variable command",
			},
			Handler:         b.HandleGetVariables,
			IsParameterized: true,
		},
		{
			Meta: tele.Command{
				Text:        "/getall",
				Description: "Get all user variables command",
			},
			Handler: b.HandleGetAllVariables,
		},
		{
			Meta: tele.Command{
				Text:        "/del",
				Description: "Delete variable command",
			},
			Handler:         b.HandleDeleteVariables,
			IsParameterized: true,
		},
		{
			Meta: tele.Command{
				Text:        "/delall",
				Description: "Delete all variables command",
			},
			Handler: b.HandleDeleteAllVariables,
		},
		{
			Meta: tele.Command{
				Text:        "/hist",
				Description: "Get solving history command",
			},
			Handler: b.HandleGetHistory,
		},
		{
			Meta: tele.Command{
				Text:        "/clear",
				Description: "Clear history and variables command",
			},
			Handler: b.HandleClearAll,
		},
		{
			Meta: tele.Command{
				Text:        "/password",
				Description: "Get current password for REST API command",
			},
			Handler: b.HandleGetPassword,
		},
		{
			Meta: tele.Command{
				Text:        "/genpassword",
				Description: "Generate new password for REST API command",
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
				Description: "Outputs detailed description",
			},
			Handler: b.HandleGreatings,
		},
	}
	return res
}
