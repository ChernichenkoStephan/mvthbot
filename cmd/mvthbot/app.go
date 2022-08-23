package main

import (
	"context"
	"time"

	"emperror.dev/errors"
	"go.uber.org/zap"

	"github.com/ChernichenkoStephan/mvthbot/internal/auth"
	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
	"github.com/ChernichenkoStephan/mvthbot/internal/fixing"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	tele "gopkg.in/telebot.v3"
)

type App struct {
	api *fiber.App

	// bot
	bot *tg.Bot

	// logging
	lg *zap.SugaredLogger

	// db
	connection *sqlx.DB

	// cache
	cache user.Cache

	// DB Wrapper
	db *user.Database

	// Repositories
	userRepository     user.UserRepository
	variableRepository user.VariableRepository
	authRepository     auth.AuthRepository
}

func InitApp(ctx context.Context /*metrics, */, lg *zap.SugaredLogger) (*App, error) {
	lg.Infoln("App setup")

	api := fiber.New(fiber.Config{
		AppName:      "Equation solving service",
		ServerHeader: "Mvthbot API",
	})

	// TODO change to normal
	cache := user.GetDummyCache(time.Hour, time.Hour)

	// DB setup
	conn, err := setupDB(lg)
	if err != nil {
		return nil, errors.Wrap(err, "Error during DB setup")
	}

	ur := user.NewUserRepository(cache, conn)
	vr := user.NewVariableRepository(cache, conn)

	ar := auth.NewAuthRepository(cache, conn)

	dblg := lg.Named("Database")
	db := user.NewDB(ur, vr, cache, ur, dblg)

	// Creating logger for bot
	blg := lg.Named("BOT")

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
	b := tg.NewBot(c, db, f, blg)

	app := &App{
		api: api,

		bot: b,

		connection: conn,

		cache: cache,

		lg: lg,

		db: db,

		userRepository:     ur,
		variableRepository: vr,
		authRepository:     ar,
	}

	// API setup
	err = setupAPI(ctx, app)
	if err != nil {
		return nil, errors.Wrap(err, "Error during API setup")
	}

	// Telegram bot setup
	err = setupBot(ctx, app, blg)
	if err != nil {
		return nil, errors.Wrap(err, "Error during Telegram bot setup")
	}

	app.lg.Infoln("App init success")
	return app, nil
}

func (app *App) Close() error {
	app.bot.Client().Stop()
	app.lg.Info("Gracefull shutdown app")

	dbErr := app.connection.Close()
	if dbErr != nil {
		app.lg.Errorf("DB connection closing fail %s", dbErr.Error())
	} else {
		app.lg.Info("DB connection closed.")
	}

	apiErr := app.api.Shutdown()
	if dbErr != nil {
		app.lg.Errorf("API connection closing fail %s", dbErr.Error())
	} else {
		app.lg.Info("API connection closed.p")
	}

	return errors.Combine(dbErr, apiErr)
}

func runApp(ctx context.Context, app *App) error {
	defer func(app *App) {
		app.lg.Info("Gracefully shuting down the App")
		app.Close()
	}(app)
	return app.api.Listen(":8080")
}
