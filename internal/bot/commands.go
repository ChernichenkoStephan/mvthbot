package bot

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) HandleAll(ctx context.Context, c tele.Context, ch chan error) {
	fmt.Println(c.Text())
	ch <- nil
}

// s command
func (b *Bot) HandleSolve(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleSolve command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// get comamnd
func (b *Bot) HandleGetVariable(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleGetVariable command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// del command
func (b *Bot) HandleDeleteVariable(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleDeleteVariable command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// delall comand
func (b *Bot) HandleDeleteAllVariables(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleDeleteAllVariables command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// hist command
func (b *Bot) HandleGetHistory(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleGetHistory command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// clear command
func (b *Bot) HandleClearAll(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleClearAll command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// password command
func (b *Bot) HandleGetPassword(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleGetPassword command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// genpassword command
func (b *Bot) HandleGeneratePassword(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleGeneratePassword command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

// help command
func (b *Bot) HandleHelp(ctx context.Context, c tele.Context, ch chan error) {
	user := c.Sender()
	resp := "HandleHelp command in process..\n"
	resp += fmt.Sprintf("With args: %v", c.Args())
	_, err := b.client.Send(user, resp)
	if err != nil {
		ch <- err
	}
	ch <- nil
}

func NewTeleHandler(f HandleFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		ch := make(chan error)
		go f(context.TODO(), c, ch)
		return <-ch
	}
}

func (b Bot) BaseCommands() *[]Command {
	res := &[]Command{
		{
			Meta: tele.Command{
				Text:        "/s",
				Description: "Solve command",
			},
			Handler: b.HandleSolve,
		},
		{
			Meta: tele.Command{
				Text:        "/get",
				Description: "Get variable command",
			},
			Handler: b.HandleGetVariable,
		},
		{
			Meta: tele.Command{
				Text:        "/del",
				Description: "Delete variable command",
			},
			Handler: b.HandleDeleteVariable,
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
	}
	return res
}
