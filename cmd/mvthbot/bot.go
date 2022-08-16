package main

import (
	"context"
	"errors"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"

	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
)

func setupBot(app *App, lg *zap.SugaredLogger) error {
	lg.Infoln("Bot setup")

	b := app.bot.Client()

	b.Use(tg.Logging(lg))
	b.Use(tg.UserCheck(app.userService))
	b.Use(tg.ArgParse)

	b.Handle(tele.OnText, func(c tele.Context) error {
		return app.bot.HandleDefault(context.TODO(), c)
	})

	commands := make([]tele.Command, 0)
	for _, cmd := range *app.bot.BaseCommands() {
		lg.Infof("Setting '%s' command", cmd.Meta.Text)

		b.Handle(cmd.Meta.Text, tg.NewTeleHandler(cmd.Handler))

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

func runBot(app *App) error {
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
