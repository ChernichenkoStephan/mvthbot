package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"emperror.dev/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	tele "gopkg.in/telebot.v3"

	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
)

func getCommandsTexts(dirPath string, lg *zap.SugaredLogger) (map[string]string, error) {
	lg.Infof("Reading command texts from: %s", dirPath)

	texts := make(map[string]string)
	const extLen = 3

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, `commands texts fetch faild`)
	}

	for _, f := range files {
		lg.Infof("Processing text from: %s", f.Name())

		if f.Name()[len(f.Name())-extLen:] != `.md` {
			return nil, fmt.Errorf("wrong file format, shuld be .txt (got: %s)", f.Name())
		}

		path := dirPath + f.Name()

		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.Wrap(err, `commands texts fetch faild`)
		}
		lg.Infof("Got %d from %s", len(contents), f.Name())

		commandName := f.Name()[:len(f.Name())-extLen]
		texts[commandName] = string(contents)
	}

	if len(texts) == 0 {
		return nil, errors.New("commands texts fetch faild")
	}
	return texts, nil
}

func setupBot(ctx context.Context, c *configuration, bot *tg.Bot, lg *zap.SugaredLogger) error {
	lg.Infoln("Bot setup")

	comandsTexts, err := getCommandsTexts(c.Bot.CommandsTextsPath, lg)
	if err != nil {
		return err
	}

	bot.ComandsTexts = comandsTexts

	b := bot.Client()

	b.Use(tg.Logging(lg))
	b.Use(bot.UserCheck())
	b.Use(tg.ArgParse)

	b.Handle(tele.OnText, func(c tele.Context) error {
		return bot.HandleDefault(ctx, c)
	})

	commands := make([]tele.Command, 0)
	for _, cmd := range *bot.BaseCommands() {
		lg.Infof("Setting '%s' command", cmd.Meta.Text)

		b.Handle(cmd.Meta.Text, tg.NewTeleHandler(ctx, cmd.Handler))

		if !cmd.IsPrivate {
			commands = append(commands, cmd.Meta)
		}
	}

	err = b.SetCommands(commands)
	if err != nil {
		lg.Errorln("Command setup failed (on set)")
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
