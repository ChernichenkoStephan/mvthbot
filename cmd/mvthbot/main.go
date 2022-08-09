package main

import (
	"context"
	"log"

	"github.com/ChernichenkoStephan/mvthbot/internal/app"
	"github.com/go-faster/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	app.Run(func(ctx context.Context /*log logger*/) error {
		log.Println("Starting...")
		g, ctx := errgroup.WithContext(ctx)

		app, err := InitApp()
		if err != nil {
			return errors.Wrap(err, "initialize")
		}

		// Run API
		g.Go(func() error {
			return runAPI(app)
		})

		// Run Bot
		g.Go(func() error {
			// return runBot(app)
			return nil
		})

		return g.Wait()
	})

}
