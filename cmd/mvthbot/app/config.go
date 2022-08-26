package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type configuration struct {
	App struct {
		ShutdownTimeout time.Duration
		WatchdogTimeout time.Duration
		Name            string
	}

	Logging struct {
		LogFileDir    string // Log path
		ErrorFileName string // Name for logs with level error
		WarnFileName  string // Name for logs with level warn
		InfoFileName  string // Name for logs with level info
		DebugFileName string // Name for logs with level debug
		MaxSize       int    // How many m of a file is greater than this number to start file segmentation
		MaxBackups    int    // MaxBackups is the maximum number of old log files to keep
		MaxAge        int    // MaxAge is the maximum number of days old log files are retained by date
	}

	Bot struct {
		PasswordLength int
		Token          string
		Greetings      string
	}

	API struct {
		// api listening port
		Port string

		// gourutine limit for api
		MaxConnections int

		// app name for api
		ServerHeader string

		// max request size
		BodyLimit int

		// in miliseconds
		ReadTimeout time.Duration

		// in miliseconds
		WriteTimeout time.Duration
	}

	Database struct {
		Driver    string // for example "postgres"
		SourceStr string // user name, db name
	}
}

func (c *configuration) defaults() {
	c.App.Name = "Mvthbot"
	c.App.ShutdownTimeout = time.Second * 5
	c.App.WatchdogTimeout = c.App.ShutdownTimeout + time.Second*5

	c.Logging.LogFileDir = exPath[:len(exPath)-3] + `logs`
	c.Logging.MaxSize = 30
	c.Logging.MaxBackups = 7
	c.Logging.MaxAge = 7

	c.Bot.PasswordLength = 8
	c.Bot.Token = os.Getenv(`BOT_TOKEN`)
	c.Bot.Greetings = "Wellcome to mvthbot"

	c.API.Port = ":8080"
	c.API.MaxConnections = 100
	c.API.ServerHeader = "Mvthbot service"
	c.API.BodyLimit = 4 * 1024 * 1024
	c.API.ReadTimeout = 1000
	c.API.WriteTimeout = 1000

	c.Database.SourceStr = os.Getenv("DATA_SOURCE_NAME")
}

func GetConfig() (*configuration, error) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var (
		c  = &configuration{}
		it interface{}
	)

	c.defaults()

	if it = viper.Get(`App.name`); it != nil {
		if s, ok := it.(string); ok {
			c.App.Name = s
		} else {
			log.Println("Wrong type for param: App.name (shuld be string)")
		}
	}

	if it = viper.Get(`App.shutdown-timeout`); it != nil {
		if intt, ok := it.(int); ok {
			c.App.ShutdownTimeout = time.Second * time.Duration(intt)
		} else {
			log.Println("Wrong type for param: App.shutdown-timeout (shuld be int)")
		}
	}

	if it = viper.Get(`App.watchdog-timeout`); it != nil {
		if intt, ok := it.(int); ok {
			c.App.WatchdogTimeout = time.Second * time.Duration(intt)
		} else {
			log.Println("Wrong type for param: App.watchdog-timeout (shuld be int)")
		}
	}

	if it = viper.Get("Log.file-dir"); it != nil {
		if s, ok := it.(string); ok {
			c.Logging.LogFileDir = s
		} else {
			log.Println("Wrong type for param: Log.file-dir (shuld be string)")
		}
	}

	if it = viper.Get("Log.error-file-name"); it != nil {
		if s, ok := it.(string); ok {
			c.Logging.ErrorFileName = s
		} else {
			log.Println("Wrong type for param: Log.error-file-name (shuld be string)")
		}
	}

	if it = viper.Get("Log.warn-file-name"); it != nil {
		if s, ok := it.(string); ok {
			c.Logging.WarnFileName = s
		} else {
			log.Println("Wrong type for param: Log.warn-file-name (shuld be string)")
		}
	}

	if it = viper.Get("Log.info-file-name"); it != nil {
		if s, ok := it.(string); ok {
			c.Logging.InfoFileName = s
		} else {
			log.Println("Wrong type for param: Log.info-file-name (shuld be string)")
		}
	}

	if it = viper.Get("Log.debug-file-name"); it != nil {
		if s, ok := it.(string); ok {
			c.Logging.DebugFileName = s
		} else {
			log.Println("Wrong type for param: Log.debug-file-name (shuld be string)")
		}
	}

	if it = viper.Get("Log.max-size"); it != nil {
		if intt, ok := it.(int); ok {
			c.Logging.MaxSize = intt
		} else {
			log.Println("Wrong type for param: Log.max-size (shuld be int)")
		}
	}

	if it = viper.Get("Log.max-backups"); it != nil {
		if intt, ok := it.(int); ok {
			c.Logging.MaxBackups = intt
		} else {
			log.Println("Wrong type for param: Log.max-backups (shuld be int)")
		}
	}

	if it = viper.Get("Log.max-age"); it != nil {
		if intt, ok := it.(int); ok {
			c.Logging.MaxAge = intt
		} else {
			log.Println("Wrong type for param: Log.max-age (shuld be int)")
		}
	}

	if it = viper.Get("Bot.password-length"); it != nil {
		if intt, ok := it.(int); ok {
			c.Bot.PasswordLength = intt
		} else {
			log.Println("Wrong type for param: Bot.password-length (shuld be int)")
		}
	}

	if it = viper.Get("Bot.greetings-text"); it != nil {
		if s, ok := it.(string); ok {
			c.Bot.Greetings = s
		} else {
			log.Println("Wrong type for param: Bot.greetings-text (shuld be string)")
		}
	}

	if it = viper.Get("Api.port"); it != nil {
		if s, ok := it.(string); ok {
			c.API.Port = s
		} else {
			log.Println("Wrong type for param: Api.port (shuld be string)")
		}
	}

	if it = viper.Get("Api.max-connections"); it != nil {
		if intt, ok := it.(int); ok {
			c.API.MaxConnections = intt
		} else {
			log.Println("Wrong type for param: Api.max-connections (shuld be int)")
		}
	}

	if it = viper.Get("Api.server-header"); it != nil {
		if s, ok := it.(string); ok {
			c.API.ServerHeader = s
		} else {
			log.Println("Wrong type for param: Api.server-header (shuld be string)")
		}
	}

	if it = viper.Get("Api.body-limit"); it != nil {
		if intt, ok := it.(int); ok {
			c.API.BodyLimit = intt
		} else {
			log.Println("Wrong type for param: Api.body-limit (shuld be int)")
		}
	}

	if it = viper.Get("Api.read-timeout"); it != nil {
		if intt, ok := it.(int); ok {
			c.API.ReadTimeout = time.Millisecond * time.Duration(intt)
		} else {
			log.Println("Wrong type for param: Api.read-timeout (shuld be int)")
		}
	}

	if it = viper.Get("Api.write-timeout"); it != nil {
		if intt, ok := it.(int); ok {
			c.API.WriteTimeout = time.Millisecond * time.Duration(intt)
		} else {
			log.Println("Wrong type for param: Api.write-timeout (shuld be int)")
		}
	}

	if it = viper.Get("Database.driver"); it != nil {
		if s, ok := it.(string); ok {
			c.Database.Driver = s
		} else {
			log.Println("Wrong type for param: Database.driver (shuld be string)")
		}
	}

	if it = viper.Get("Database.source-str"); it != nil {
		if s, ok := it.(string); ok {
			c.Database.SourceStr = s
		} else {
			log.Println("Wrong type for param: Database.source (shuld be string)")
		}
	}
	if c.Bot.Token == `` {
		return nil, errors.New("NEED TELEGRAM BOT TOKEN IN ENV VARS")
	}
	if os.Getenv("SECRET") == `` {
		return nil, errors.New("NEED SECRET FOR PASSWORD GEN")
	}

	return c, nil
}

func viperSetup() {
	log.Println("viperSetup...")

	if cfgFile != `` {

		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {

		// Look in default directories
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./configs/")
		viper.AddConfigPath("../configs/")
		viper.AddConfigPath(exPath)
	}

	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
}
