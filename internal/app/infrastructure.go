package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	exitCodeOk             = 0
	exitCodeApplicationErr = 1
	exitCodeWatchdog       = 1
)

const (
	shutdownTimeout = time.Second * 5
	watchdogTimeout = shutdownTimeout + time.Second*5
)

const EnvLogLevel = "LOG_LEVEL"

func Run(f func(ctx context.Context /*log logger*/) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		if err := f(ctx /*log logger*/); err != nil {
			return err
		}
		return nil
	})
	go func() {
		// Guaranteed way to kill application.
		<-ctx.Done()

		// Context is canceled, giving application time to shut down gracefully.
		//lg.Info("Waiting for application shutdown")
		fmt.Printf("\nCaiting for application shutdown\n")
		time.Sleep(watchdogTimeout)

		// Probably deadlock, forcing shutdown.
		//lg.Warn("Graceful shutdown watchdog triggered: forcing shutdown")
		fmt.Printf("\nGraceful shutdown watchdog triggered: forcing shutdown\n")
		os.Exit(exitCodeWatchdog)
	}()

	// Note that we are calling os.Exit() here and no
	if err := wg.Wait(); err != nil {
		//lg.Error("Failed",
		//	zap.Error(err),
		//)
		fmt.Printf("Failed %v", err)
		os.Exit(exitCodeApplicationErr)
	}

	os.Exit(exitCodeOk)
}
