package main

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
	defer func() {
		if panicked {
			app.lg.Errorln("telegram bot runing error")
		}
		ctx.Done()
	}()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		app.bot.Client().Start()
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		app.bot.Client().Stop()
		app.lg.Info("Gracefull shutdown app")
		return nil
	})

	panicked = false

	return group.Wait()
}
