package app

import (
	"context"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	exitCodeOk             = 0
	exitCodeApplicationErr = 1
	exitCodeWatchdog       = 1
)

type AppConfig struct {
	ShutdownTimeout time.Duration
	WatchdogTimeout time.Duration
	Logger          *zap.SugaredLogger
}

func Run(c *AppConfig, f func(ctx context.Context, lg *zap.SugaredLogger) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		if err := f(ctx, c.Logger); err != nil {
			return err
		}
		return nil
	})

	go func() {
		// Guaranteed way to kill application.
		<-ctx.Done()

		// Context is canceled, giving application time to shut down gracefully.
		c.Logger.Infof("\nCaiting for application shutdown\n")
		time.Sleep(c.WatchdogTimeout)

		// Probably deadlock, forcing shutdown.
		c.Logger.Infof("\nGraceful shutdown watchdog triggered: forcing shutdown\n")
		os.Exit(exitCodeWatchdog)
	}()

	// Note that we are calling os.Exit() here and no
	if err := wg.Wait(); err != nil {
		c.Logger.Errorf("Failed %v", err)
		os.Exit(exitCodeApplicationErr)
	}

	os.Exit(exitCodeOk)
}
