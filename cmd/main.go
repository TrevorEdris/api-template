package main

import (
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/TrevorEdris/api-template/app"
	"github.com/TrevorEdris/api-template/somedatabase"
)

type configuration struct {
	// Defined at the top-level
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	// Defined by inner packages
	ServerConfig app.ServerConfig
}

func main() {
	var cfg configuration

	logCtx := log15.Ctx{
		"module": app.AppName,
	}
	log := log15.New(logCtx)

	// Load the environment variables if specified in .env
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Crit("Unable to load env file", "error", err)
		os.Exit(1)
	}

	// Parse the environment variables into a configuration struct
	err = envconfig.Process(app.EnvPrefix, &cfg)
	if err != nil {
		log.Crit("Unable to process config", "error", err)
		os.Exit(1)
	}

	// Initialize the base logger
	logLevel, err := log15.LvlFromString(strings.ToLower(cfg.LogLevel))
	if err != nil {
		log15.Warn("Improper log level specified; defaulting to 'info'", "error", err)
		logLevel = log15.LvlInfo
	}
	logHandler := log15.LvlFilterHandler(logLevel, log15.StreamHandler(os.Stdout, log15.JsonFormat()))
	log.SetHandler(logHandler)

	// Create a connection to some database
	someDBClient := somedatabase.NewSomeDatabaseClient("some_config")

	// Create the app
	application := app.New(cfg.ServerConfig, log, &someDBClient)
	application.Run()
}
