package main

import (
	"log"

	"github.com/go-faster/errors"

	"github.com/ChernichenkoStephan/mvthbot/internal/auth"
	tg "github.com/ChernichenkoStephan/mvthbot/internal/bot"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/gofiber/fiber/v2"
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
	authRepository     auth.UserAuthRepository
}

func InitApp( /*metrics logger db*/ ) (*App, error) {
	log.Println("App setup")

	api := fiber.New(fiber.Config{
		AppName:      "Equation solving service",
		ServerHeader: "Fiber",
	})

	ur := user.NewImdbUserRepository()
	vr := user.NewImdbVariableRepository()
	ar := auth.NewUserAuthRepository()

	us := user.NewUserService(ur)
	vs := user.NewVariableService(vr)

	b := tg.NewBot()

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
	err := setupDB(app)
	if err != nil {
		return nil, errors.Wrap(err, "DB setup")
	}

	// API setup
	err = setupAPI(app)
	if err != nil {
		return nil, errors.Wrap(err, "API setup")
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
