package main

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

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

func InitApp(ctx context.Context /*metrics, */, c *configuration, lg *zap.SugaredLogger) (*App, error) {
	lg.Infoln("App setup")

	api := fiber.New(fiber.Config{
		AppName:      c.App.Name,
		ServerHeader: c.API.ServerHeader,
		BodyLimit:    c.API.BodyLimit,
		ReadTimeout:  c.API.ReadTimeout,
		WriteTimeout: c.API.WriteTimeout,
	})

	cache := user.GetCache(time.Hour, time.Hour)

	// DB setup
	conn, err := setupDB(c, lg)
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

	pref := tele.Settings{
		Token:  c.Bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c tele.Context) {
			blg.Errorf("Got error in Telegram bot: %v", err)
		},
	}

	tgc, err := tele.NewBot(pref)
	if err != nil {
		return nil, errors.Wrap(err, "Error during Telegram bot setup")
	}

	f := fixing.New()

	b := tg.NewBot(tgc, db, f, blg, &tg.BotConfig{
		PasswordLength: c.Bot.PasswordLength,
		GreetText:      c.Bot.Greetings,
	})

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
	err = setupAPI(ctx, c, app)
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

func (app *App) Close() {

	dbErr := app.connection.Close()
	if dbErr != nil {
		app.lg.Errorf("DB connection closing fail %s", dbErr.Error())
	} else {
		app.lg.Info("DB connection closed.")
	}

	logErr := app.lg.Sync()
	if logErr != nil {
		fmt.Println(logErr)
	}

}

func runApp(ctx context.Context, port string, app *App) error {
	defer func(app *App) {
		app.lg.Info("Gracefully shuting down the App")
		app.Close()
	}(app)

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return app.api.Listen(port)
	})

	group.Go(func() error {
		<-ctx.Done()

		app.lg.Info("Shuting down API server")
		apiErr := app.api.Shutdown()
		if apiErr != nil {
			app.lg.Errorf("API connection closing fail %s", apiErr.Error())
		} else {
			app.lg.Info("API connection closed.")
		}

		return nil
	})

	return group.Wait()
}
