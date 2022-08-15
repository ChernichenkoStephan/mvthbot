package main

import (
	"errors"
	"fmt"

	"github.com/ChernichenkoStephan/mvthbot/internal/auth"
	"github.com/ChernichenkoStephan/mvthbot/internal/misc"
	"github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func setupAPI(app *App) error {
	app.lg.Infoln("Api setup")

	app.api.Use(cors.New())
	// api.Use(etag.New())
	app.api.Use(favicon.New())
	app.api.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
			return errors.New("Requests overload")
		},
	}))

	// Making logger for API view
	lg := app.lg.Named("API")

	// Logging with zap
	app.api.Use(func(c *fiber.Ctx) error {

		lg.Infof("[%s]:%s | %s | %s", c.IP(), c.Port(), c.Method(), c.Path())
		return c.Next()
	})

	app.api.Use(recover.New())
	app.api.Use(requestid.New())

	// Prepare our endpoints for the API.
	misc.NewMiscHandler(app.api.Group("/api/v1"))
	solving.NewSolveHandler(app.api.Group("/api/v1/solve"))
	auth.NewAuthHandler(app.api.Group("/api/v1/auth"), app.authRepository, app.bot, lg)
	user.NewVariableHandler(
		app.api.Group("/api/v1/variables", auth.JWTMiddleware()),
		app.userService,
		app.veriableService,
	)
	user.NewHistoryHandler(
		app.api.Group("/api/v1/history", auth.JWTMiddleware()),
		app.userService,
	)

	// Prepare an endpoint for 'Not Found'.
	app.api.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})

		return fmt.Errorf("Unknown route %s", c.OriginalURL())
	})

	app.lg.Infoln("Api setup success")
	return nil
}

func runAPI(app *App) error {
	return app.api.Listen(":8080")
}
