package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/ChernichenkoStephan/mvthbot/internal/utils"

	tele "gopkg.in/telebot.v3"
)

func NewTeleHandler(f HandleFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		ch := make(chan error)
		go f(context.TODO(), c, ch)
		err := <-ch
		log.Println(err)
		return err
	}
}

func getDest(c tele.Context) tele.Recipient {
	switch t := c.Chat().Type; t {
	case "private":
		return c.Sender()
	default:
		log.Printf("[receaved] ChatType: %s", t)
	}
	return nil
}

func (b *Bot) HandleDefault(ctx context.Context, c tele.Context, ch chan error) {
	dest := c.Sender()

	// To work with empty (example: 2+2) statements only in bot
	if c.Chat().Type == "private" {

		ctx := context.TODO()
		resp, err := b.process(ctx, dest.ID, c.Get("statements"))
		if err != nil {
			resp = err.Error()
		}

		t := err
		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- fmt.Errorf("Got: %v %v", err, t)
		}
	}

	ch <- nil
}

// start command
func (b *Bot) HandleGreatings(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	_, err := b.client.Send(user, "Hi, let's go!")
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// s command
func (b *Bot) HandleSolve(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)

	ctx = context.TODO()
	resp, err := b.process(ctx, c.Sender().ID, c.Get("statements"))
	if err != nil {
		resp = err.Error()
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("Got: %v %v", err, t)
	}

	ch <- nil
}

// get comamnd
func (b *Bot) HandleGetVariables(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string

	ctx = context.TODO()
	vs, err := b.variablesService.GetWithNames(ctx, c.Sender().ID, c.Args())
	if err != nil {
		resp = err.Error()
	} else {
		builder := NewOutputBuilder()
		for n, v := range vs {
			builder.WriteVariable(n)
			builder.WriteValue(v)
			builder.LineBreak()
		}
		resp = builder.String()
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// getall comamnd
func (b *Bot) HandleGetAllVariables(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string

	ctx = context.TODO()
	vs, err := b.variablesService.GetAll(ctx, c.Sender().ID)
	if err != nil {
		resp = err.Error()
	} else {
		builder := NewOutputBuilder()
		for n, v := range vs {
			builder.WriteVariable(n)
			builder.WriteValue(v)
			builder.LineBreak()
		}
		resp = builder.String()
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// del command
func (b *Bot) HandleDeleteVariables(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string

	ctx = context.TODO()
	err := b.variablesService.DeleteWithNames(ctx, c.Sender().ID, c.Args())
	if err != nil {
		resp = err.Error()
	} else {
		resp = "Success"
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// delall comand
func (b *Bot) HandleDeleteAllVariables(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string

	ctx = context.TODO()
	err := b.variablesService.DeleteAll(ctx, c.Sender().ID)
	if err != nil {
		resp = err.Error()
	} else {
		resp = "Success"
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// hist command
func (b *Bot) HandleGetHistory(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string

	ctx = context.TODO()
	sts, err := b.userService.GetHistory(ctx, c.Sender().ID)
	if err != nil {
		resp = err.Error()
	} else {
		log.Printf("%v", *sts)
		builder := NewOutputBuilder()
		for _, s := range *sts {
			builder.WriteFull(&s)
		}
		resp = builder.String()
		if resp == "" {
			resp = "empty"
		}
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// clear command
func (b *Bot) HandleClearAll(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string

	ctx = context.TODO()
	err := b.userService.Clear(ctx, c.Sender().ID)
	if err != nil {
		resp = err.Error()
	} else {
		resp = "Success"
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// password command
func (b *Bot) HandleGetPassword(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string
	var err error

	if c.Chat().Type == "private" {

		ctx = context.TODO()
		u, err := b.userService.Get(ctx, c.Sender().ID)
		if err != nil {
			resp = err.Error()
		} else {
			if u.Password == "" {
				resp = "No password. Generete one with /genpassword command"
			} else {
				resp = u.Password
			}
		}

	} else {
		resp = "Command forbiden, use only in private bot chat"
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// genpassword command
func (b *Bot) HandleGeneratePassword(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)
	var resp string
	var err error

	if c.Chat().Type == "private" {

		ctx = context.TODO()
		u, err := b.userService.Get(ctx, c.Sender().ID)
		if err != nil {
			resp = err.Error()
		} else {
			// TODO fix to config
			u.Password = utils.GenPassword(8)
			resp = fmt.Sprintf("%v", u.Password)
		}

	} else {
		resp = "Command forbiden, use only in private bot chat"
	}

	t := err
	_, err = b.client.Send(dest, resp)
	if err != nil {
		ch <- fmt.Errorf("%s, %s", t, err)
	}
	ch <- nil
}

// help command
func (b *Bot) HandleHelp(ctx context.Context, c tele.Context, ch chan error) {
	dest := getDest(c)

	resp := "HandleHelp command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(dest, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
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
