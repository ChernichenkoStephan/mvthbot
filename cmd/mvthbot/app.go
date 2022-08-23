package main

import (
	"time"

	"emperror.dev/errors"
	"go.uber.org/zap"

	"github.com/ChernichenkoStephan/mvthbot/internal/auth"
	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	api *fiber.App

	// bot
	bot *tg.Bot

	// logging
	lg *zap.SugaredLogger

	// db

	// cache
	cache user.Cache

	// Services
	db *user.Database

	// Repositories
	userRepository     user.UserRepository
	variableRepository user.VariableRepository
	authRepository     auth.AuthRepository
}

func InitApp( /*metrics db*/ lg *zap.SugaredLogger) (*App, error) {
	lg.Infoln("App setup")

	api := fiber.New(fiber.Config{
		AppName:      "Equation solving service",
		ServerHeader: "Mvthbot API",
	})

	// TODO change to normal
	cache := user.GetCache(time.Hour, time.Hour)

	// DB setup
	conn, err := setupDB(lg)
	if err != nil {
		return nil, errors.Wrap(err, "Error during DB setup")
	}

	ur := user.NewUserRepository(cache, conn)
	vr := user.NewVariableRepository(cache, conn)

	ar := auth.NewAuthRepository(cache, conn)

	dblg := lg.Named("Database")
	db := user.NewDB(ur, vr, cache, dblg)

	// Creating logger for bot
	//blg := lg.Named("BOT")

	/*
		// TODO Token from Viper
		pref := tele.Settings{
			Token:  "5597673919:AAGdW5TVuWkFkvCf87knskCPlg7HUoipSTY",
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
			OnError: func(err error, c tele.Context) {
				blg.Errorf("Got error in Telegram bot: %v", err)
			},
		}

		c, err := tele.NewBot(pref)
		if err != nil {
			return nil, errors.Wrap(err, "Error during Telegram bot setup")
		}

		f := fixing.New()
		b := tg.NewBot(c, us, vs, f, blg)
	*/

	app := &App{
		api: api,

		//bot: b,

		cache: cache,

		lg: lg,

		db: db,

		userRepository:     ur,
		variableRepository: vr,
		authRepository:     ar,
	}

	// API setup
	err = setupAPI(app)
	if err != nil {
		return nil, errors.Wrap(err, "Error during API setup")
	}

	/*
		// Telegram bot setup
		err = setupBot(app, blg)
		if err != nil {
			return nil, errors.Wrap(err, "Error during Telegram bot setup")
		}
	*/

	app.lg.Infoln("App init success")
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
