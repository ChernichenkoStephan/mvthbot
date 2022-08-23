package main

import (
	"context"
	"errors"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"

	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
)

func setupBot(ctx context.Context, app *App, lg *zap.SugaredLogger) error {
	lg.Infoln("Bot setup")

	b := app.bot.Client()

	b.Use(tg.Logging(lg))
	b.Use(tg.UserCheck(app.db))
	b.Use(tg.ArgParse)

	b.Handle(tele.OnText, func(c tele.Context) error {
		return app.bot.HandleDefault(ctx, c)
	})

	commands := make([]tele.Command, 0)
	for _, cmd := range *app.bot.BaseCommands() {
		lg.Infof("Setting '%s' command", cmd.Meta.Text)

		b.Handle(cmd.Meta.Text, tg.NewTeleHandler(ctx, cmd.Handler))

		if !cmd.IsParameterized {
			commands = append(commands, cmd.Meta)
		}
	}

	err := b.SetCommands(commands)
	if err != nil {
		lg.Errorln("Command setup failed")
	}

	lg.Infoln("Bot setup success")
	return nil

}

func runBot(ctx context.Context, app *App) error {
	panicked := true
	var err error
	defer func() {
		if panicked {
			err = errors.New("Telegram bot runing error")
		}
	}()
	app.bot.Client().Start()
	panicked = false
	return err
}
