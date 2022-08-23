package main

import (
	"context"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/app"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"go.uber.org/zap"
)

func main() {

	app.Run(func(ctx context.Context, lg *zap.SugaredLogger) error {
		lg.Infoln("Starting...")
		// g, ctx := errgroup.WithContext(ctx)

		app, err := InitApp(lg)
		if err != nil {
			return errors.Wrap(err, "App init failed")
		}

		return user.DummyRun(app.db, user.DUMMY_USER_RUN)

		/*
				// Run API
				g.Go(func() error {
					return runAPI(app)
				})

				// Run Bot
				g.Go(func() error {
					return runBot(app)
				})

			return g.Wait()
		*/
	})

}
