package main

import (
	"context"
	"fmt"
	"log"

	tele "gopkg.in/telebot.v3"

	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
)

func setupBot(app *App) error {
	log.Println("Bot setup")

	b := app.bot.Client()

	b.Use(tg.ArgParse)

	b.Handle(tele.OnText, func(c tele.Context) error {
		ch := make(chan error)
		go app.bot.HandleAll(context.TODO(), c, ch)
		return <-ch
	})

	commands := make([]tele.Command, len(*app.bot.BaseCommands()))
	for i, cmd := range *app.bot.BaseCommands() {
		log.Printf("Setting '%s' command\n", cmd.Meta.Text)
		b.Handle(cmd.Meta.Text, tg.NewTeleHandler(cmd.Handler))
		commands[i] = cmd.Meta
	}

	return b.SetCommands(commands)
}

func runBot(app *App) error {
	panicked := true
	var err error
	defer func() {
		if panicked {
			err = fmt.Errorf("Telegram bot runing error")
		}
	}()
	app.bot.Client().Start()
	panicked = false
	return err
}
