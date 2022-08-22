package main

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/app"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func step(e error, opName string, res interface{}) bool {
	if e != nil {
		fmt.Println(e)
		return false
	} else {
		fmt.Println(opName, " SUCCESS")
	}
	fmt.Println(res)
	fmt.Println("next?")
	o, err := fmt.Scanln()
	if err != nil {
		fmt.Println(err)
		return false
	} else if o != 0 {
		return false
	}
	fmt.Println(o)
	return true
}

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
