package main

import (
	"log"
	"time"

	"github.com/go-faster/errors"

	"github.com/ChernichenkoStephan/mvthbot/internal/auth"
	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
	"github.com/ChernichenkoStephan/mvthbot/internal/fixing"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/gofiber/fiber/v2"

	tele "gopkg.in/telebot.v3"
)

type App struct {
	api *fiber.App

	// bot
	bot *tg.Bot

	// db

	// Services
	userService     user.UserService
	veriableService user.VariableService

	// Repositories
	userRepository     user.UserRepository
	variableRepository user.VariableRepository
	authRepository     auth.AuthRepository
}

func InitApp( /*metrics logger db*/ ) (*App, error) {
	log.Println("App setup")

	api := fiber.New(fiber.Config{
		AppName:      "Equation solving service",
		ServerHeader: "Fiber",
	})

	// TODO change to normal
	cache := user.GetTestCache(time.Hour, time.Hour)

	ur := user.NewUserRepository(cache)
	vr := user.NewVariableRepository(cache)
	ar := auth.NewAuthRepository(cache)

	us := user.NewUserService(ur)
	vs := user.NewVariableService(vr)

	// TODO Token from Viper
	pref := tele.Settings{
		Token:  "5597673919:AAGdW5TVuWkFkvCf87knskCPlg7HUoipSTY",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c tele.Context) {
			log.Printf("Got error: %v", err)
		},
	}

	c, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "Error during Telegram bot setup")
	}

	f := fixing.New()
	b := tg.NewBot(c, us, vs, f)

	app := &App{
		api: api,

		bot: b,

		userService:     us,
		veriableService: vs,

		userRepository:     ur,
		variableRepository: vr,
		authRepository:     ar,
	}

	// DB setup
	err = setupDB(app)
	if err != nil {
		return nil, errors.Wrap(err, "Error during DB setup")
	}

	// API setup
	err = setupAPI(app)
	if err != nil {
		return nil, errors.Wrap(err, "Error during API setup")
	}

	// Telegram bot setup
	err = setupBot(app)
	if err != nil {
		return nil, errors.Wrap(err, "Error during Telegram bot setup")
	}

	return app, nil
}

func (b *App) Close() error {
	//err := multierr.Append(b.stateStorage.db.Close(), b.db.Close())
	//if b.index != nil {
	//	err = multierr.Append(err, b.index.Close())
	//}
	//return err
	return nil
}
