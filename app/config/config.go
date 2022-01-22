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
		HTTP
		Metrics
	}

	// App defines the configs needed for the application itself.
	App struct {
		Name        string        `env:"APP_NAME,default=backfill"`
		Environment Environment   `env:"APP_ENVIRONMENT,default=local"`
		LogLevel    LogLevel      `env:"LOG_LEVEL,default=info"`
		Timeout     time.Duration `env:"APP_TIMEOUT,default=20s"`
		Storage     Storage       `env:"APP_STORAGE,default=local"`
	}

	// HTTP stores the configuration for the HTTP server.
	HTTP struct {
		Hostname     string        `env:"HTTP_HOSTNAME,default=0.0.0.0"`
		Port         uint16        `env:"HTTP_PORT,default=8000"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
		IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
		TLS          struct {
			Enabled     bool   `env:"HTTP_TLS_ENABLED,default=false"`
			Certificate string `env:"HTTP_TLS_CERTIFICATE"`
			Key         string `env:"HTTP_TLS_KEY"`
		}
	}

	// Metrics defines the configs needed for collecting metrics.
	Metrics struct {
		Enabled bool   `env:"METRICS_ENABLED,default=false"`
		Addr    string `env:"METRICS_ADDRESS"`
		BufLen  int    `env:"METRICS_BUFFER,default=5"`
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
