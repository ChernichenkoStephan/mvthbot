package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/ChernichenkoStephan/mvthbot/internal/logging"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	exitCodeOk             = 0
	exitCodeApplicationErr = 1
	exitCodeWatchdog       = 1
)

const (
	shutdownTimeout = time.Second * 1
	watchdogTimeout = shutdownTimeout + time.Second*1
)

const EnvLogLevel = "LOG_LEVEL"

func Run(f func(ctx context.Context, lg *zap.SugaredLogger) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	opt := &logging.Options{
		LogFileDir: exPath + "/logs",
		AppName:    "logtool",
		MaxSize:    30,
		MaxBackups: 7,
		MaxAge:     7,
		Config:     zap.Config{},
	}
	opt.Development = true

	logging.InitLogger(opt)
	lg := logging.GetLogger()

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		if err := f(ctx, lg.SugaredLogger); err != nil {
			return err
		}
		return nil
	})
	go func() {
		// Guaranteed way to kill application.
		<-ctx.Done()

		// Context is canceled, giving application time to shut down gracefully.
		lg.Infof("\nCaiting for application shutdown\n")
		time.Sleep(watchdogTimeout)

		// Probably deadlock, forcing shutdown.
		lg.Infof("\nGraceful shutdown watchdog triggered: forcing shutdown\n")
		os.Exit(exitCodeWatchdog)
	}()

	// Note that we are calling os.Exit() here and no
	if err := wg.Wait(); err != nil {
		lg.Errorf("Failed %v", err)
		os.Exit(exitCodeApplicationErr)
	}

	os.Exit(exitCodeOk)
}
