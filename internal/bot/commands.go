package bot

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"
	"go.uber.org/zap"

	tele "gopkg.in/telebot.v3"
)

func NewTeleHandler(f HandleFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		ch := make(chan error, 2)
		go f(context.TODO(), c, ch)
		err := <-ch
		return err
	}
}

func getDest(c tele.Context, lg *zap.SugaredLogger) (tele.Recipient, error) {
	switch t := c.Chat().Type; t {
	case "private":
		return c.Sender(), nil
	default:
		//lg.Warnf("[receaved] ChatType: %s", t)
	}
	return nil, fmt.Errorf("Forbiden chat type %s", c.Chat().Type)
}

func (b *Bot) HandleDefault(ctx context.Context, c tele.Context, ch chan error) {
	dest := c.Sender()

	// To work with empty (example: 2+2) statements only in bot
	if c.Chat().Type == "private" {

		ctx := context.TODO()
		resp, err := b.process(ctx, dest.ID, c.Get("statements"))
		if err != nil {
			resp = err.Error()
			ch <- errors.Wrap(err, "Statements processing failed")
		}

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}
	}

	ch <- nil
	close(ch)
}

// start command
func (b *Bot) HandleGreatings(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	_, err := b.client.Send(user, "Hi, let's go!")
	if err != nil {
		ch <- errors.Wrap(err, "Reply failed")
	}
	ch <- nil
}

// s command
func (b *Bot) HandleSolve(ctx context.Context, c tele.Context, ch chan error) {

	if dest, err := getDest(c, b.logger); err == nil {

		ctx = context.TODO()
		resp, err := b.process(ctx, c.Sender().ID, c.Get("statements"))
		if err != nil {
			resp = err.Error()
			ch <- errors.Wrap(err, "Statements processing failed")
		}

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

		ch <- nil

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)

}

// get comamnd
func (b *Bot) HandleGetVariables(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		ctx = context.TODO()
		vs, err := b.variablesService.GetWithNames(ctx, c.Sender().ID, c.Args())
		if err != nil {
			ch <- errors.Wrap(err, "Error during getting user variables")
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

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

		ch <- nil

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// getall comamnd
func (b *Bot) HandleGetAllVariables(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		ctx = context.TODO()
		vs, err := b.variablesService.GetAll(ctx, c.Sender().ID)
		if err != nil {
			ch <- errors.Wrap(err, "Error during getting user variables")
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

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

		ch <- nil

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// del command
func (b *Bot) HandleDeleteVariables(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		ctx = context.TODO()
		err := b.variablesService.DeleteWithNames(ctx, c.Sender().ID, c.Args())
		if err != nil {
			ch <- errors.Wrap(err, "Error during deleting user variables")
			resp = err.Error()
		} else {
			resp = "Success"
		}

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

		ch <- nil

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// delall comand
func (b *Bot) HandleDeleteAllVariables(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		ctx = context.TODO()
		err := b.variablesService.DeleteAll(ctx, c.Sender().ID)
		if err != nil {
			ch <- errors.Wrap(err, "Error during deleting user variables")
			resp = err.Error()
		} else {
			resp = "Success"
		}

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// hist command
func (b *Bot) HandleGetHistory(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		ctx = context.TODO()
		sts, err := b.userService.GetHistory(ctx, c.Sender().ID)
		if err != nil {
			ch <- errors.Wrap(err, "Error during deleting user history")
			resp = err.Error()
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

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// clear command
func (b *Bot) HandleClearAll(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		ctx = context.TODO()
		err := b.userService.Clear(ctx, c.Sender().ID)
		if err != nil {
			ch <- errors.Wrap(err, "Error during clear user data")
			resp = err.Error()
		} else {
			resp = "Success"
		}

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// password command
func (b *Bot) HandleGetPassword(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		if c.Chat().Type == "private" {

			ctx = context.TODO()
			u, err := b.userService.Get(ctx, c.Sender().ID)
			if err != nil {
				ch <- errors.Wrap(err, "Error during geting user")
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

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// genpassword command
func (b *Bot) HandleGeneratePassword(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {
		var resp string

		if c.Chat().Type == "private" {

			ctx = context.TODO()
			u, err := b.userService.Get(ctx, c.Sender().ID)
			if err != nil {
				ch <- errors.Wrap(err, "Error during geting user")
				resp = err.Error()
			} else {
				// TODO fix to config
				u.Password = utils.GenPassword(8)
				resp = fmt.Sprintf("%v", u.Password)
			}

		} else {
			resp = "Command forbiden, use only in private bot chat"
		}

		_, err = b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
}

// help command
func (b *Bot) HandleHelp(ctx context.Context, c tele.Context, ch chan error) {
	if dest, err := getDest(c, b.logger); err == nil {

		resp := "HandleHelp command in process..\n"
		resp += fmt.Sprintf("With args: %v", c.Args())

		_, err := b.client.Send(dest, resp)
		if err != nil {
			ch <- errors.Wrap(err, "Bot reply failed")
		}

	} else {
		ch <- errors.Wrap(err, "Get destination fail")
	}
	close(ch)
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
