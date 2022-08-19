package main

import (
	"context"
	"fmt"

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

		//err = app.variableRepository.Delete(context.TODO(), 11111, `a`)
		//err = app.variableRepository.DeleteWithNames(context.TODO(), 11111, []string{`a`, `b`})
		err = app.variableRepository.DeleteAll(context.TODO(), 11111)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("SUCCESS")
		}
		//fmt.Printf("%#v\n", v)

		/*
			// Run API
			g.Go(func() error {
				return runAPI(app)
			})

			// Run Bot
			g.Go(func() error {
				return runBot(app)
			})
		*/

		return g.Wait()
	})

}
