package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/app"
	"github.com/ChernichenkoStephan/mvthbot/internal/logging"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	logsDir string
	cfgFile string
	exPath  string
	token   string
	port    string

	rootCmd = &cobra.Command{
		Use:   "cmw-statistics-cla",
		Short: "cmw-statistics-cla: A sample CLI application for fetching stats form Telegram Chats and Channels",
		Long:  `cmw-statistics-cla: A sample CLI application written in Go for fetching stats form Telegram Chats and Channels`,
		Run: func(cmd *cobra.Command, args []string) {

			conf, err := GetConfig()
			if err != nil {
				log.Panicf("Error during config read: %s", err.Error())
			}

			opt := &logging.Options{
				LogFileDir: exPath[:len(exPath)-3] + `logs`,
				AppName:    "logtool",
				MaxSize:    30,
				MaxBackups: 7,
				MaxAge:     7,
				Config:     zap.Config{},
			}
			opt.Development = true

			logging.InitLogger(opt)
			lg := logging.GetLogger()
			slg := lg.SugaredLogger

			appConf := &app.AppConfig{
				ShutdownTimeout: conf.App.ShutdownTimeout,
				WatchdogTimeout: conf.App.WatchdogTimeout,
				Logger:          slg,
			}

			app.Run(appConf, func(ctx context.Context, lg *zap.SugaredLogger) error {
				lg.Infoln("Starting...")
				g, ctx := errgroup.WithContext(ctx)

				app, err := InitApp(ctx, conf, slg)
				if err != nil {
					return errors.Wrap(err, "App init failed")
				}

				// Run Bot
				g.Go(func() error {
					return runBot(ctx, app)
				})

				// Run API and DB
				g.Go(func() error {
					return runApp(ctx, conf.API.Port, app)
				})

				return g.Wait()
			})
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath = filepath.Dir(ex)
	log.Printf("Executing in %v", exPath)

	rootCmd.PersistentFlags().StringVarP(&logsDir, `log`, `l`, exPath[:len(exPath)-3]+`logs`, `log dir path`)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, `config`, `c`, ``, `config file (default is ./configs/config.yaml)`)
	rootCmd.PersistentFlags().StringVarP(&token, `token`, `t`, os.Getenv("BOT_TOKEN"), `Telegram bot token`)
	rootCmd.PersistentFlags().StringVarP(&port, `port`, `p`, `8080`, `API server listening port`)

	rootCmd.SetHelpTemplate(`Available flags: "--config" (-c), "--configdir" (-d), "--token" (-t), "--port" (-p)`)

	cobra.OnInitialize(viperSetup)
}

func main() {
	Execute()
}
