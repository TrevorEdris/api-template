package config

import (
	"os"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type (
	Environment string
	LogLevel    string
	Storage     string
)

const (
	EnvLocal   Environment = "local"
	EnvTest    Environment = "test"
	EnvDev     Environment = "dev"
	EnvStaging Environment = "staging"
	EnvLT      Environment = "lt"
	EnvQA      Environment = "qa"
	EnvProd    Environment = "prod"

	StorageLocal    Storage = "local"
	StorageDynamoDB Storage = "dynamodb"

	LvlDbg  LogLevel = "debug"
	LvlInfo LogLevel = "info"
	LvlWarn LogLevel = "warn"
	LvlErr  LogLevel = "error"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in. This must be called prior to loading the configuration in order for it
// to take effect.
func SwitchEnvironment(env Environment) {
	if err := os.Setenv("APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config is the aggregation of all necessary configurations.
	Config struct {
		App
	}

	// App defines the configs needed for the application itself.
	App struct {
		Name        string        `env:"APP_NAME,default=backfill"`
		Environment Environment   `env:"APP_ENVIRONMENT,default=local"`
		LogLevel    LogLevel      `env:"LOG_LEVEL,default=info"`
		Timeout     time.Duration `env:"APP_TIMEOUT,default=20s"`
		Storage     Storage       `env:"APP_STORAGE,default=local"`
	}
)

// New loads the configuration based on the environment variables.
func New() (Config, error) {
	var cfg Config
	err := godotenv.Load()

	// If a .env file exists but was unable to be loaded
	if err != nil && !os.IsNotExist(err) {
		return Config{}, err
	}

	err = envdecode.StrictDecode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
