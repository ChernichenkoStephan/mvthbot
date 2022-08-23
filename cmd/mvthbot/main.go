package main

import (
	"context"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/app"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {

	app.Run(func(ctx context.Context, lg *zap.SugaredLogger) error {
		lg.Infoln("Starting...")
		g, ctx := errgroup.WithContext(ctx)

		app, err := InitApp(ctx, lg)
		if err != nil {
			return errors.Wrap(err, "App init failed")
		}

		// Run Bot
		g.Go(func() error {
			return runBot(ctx, app)
		})

		// Run API and DB
		g.Go(func() error {
			return runApp(ctx, app)
		})

		return g.Wait()
	})

}
