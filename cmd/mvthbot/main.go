package main

import (
	"context"

	"github.com/ChernichenkoStephan/mvthbot/internal/app"
	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {

	app.Run(func(ctx context.Context, lg *zap.SugaredLogger) error {
		lg.Infoln("Starting...")
		g, ctx := errgroup.WithContext(ctx)

		app, err := InitApp(lg)
		if err != nil {
			return errors.Wrap(err, "App init failed")
		}

		// Run API
		g.Go(func() error {
			return runAPI(app)
		})

		// Run Bot
		g.Go(func() error {
			return runBot(app)
		})

		return g.Wait()
	})

}
