package config

import (
<<<<<<< HEAD
	"context"
	"errors"
=======
>>>>>>> 11540d0... Refactor to Controller/Service/Repo pattern
	"os"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		General
		Metrics
		API
		RDS
	}

	General struct {
		App      string        `env:"APP"`
		Name     string        `env:"APP_NAME"`
		LogLevel string        `env:"APP_LOG_LEVEL"`
		Timeout  time.Duration `env:"APP_TIMEOUT"`
	}

	Metrics struct {
		URL            string `env:"METRICS_URL"`
		BufLen         int    `env:"METRICS_BUF_LEN"`
		Enabled        bool   `env:"METRICS_ENABLED"`
		TracingEnabled bool   `env:"METRICS_TRACING_ENABLED"`
	}

	API struct {
		Host         string        `env:"API_HOST"`
		Port         int           `env:"API_PORT"`
		ReadTimeout  time.Duration `env:"API_READ_TIMEOUT"`
		WriteTimeout time.Duration `env:"API_WRITE_TIMEOUT"`
	}

	RDS struct {
		Host            string        `env:"RDS_HOST"`
		Username        string        `env:"RDS_USERNAME"`
		Password        string        `env:"RDS_PASSWORD"`
		Database        string        `env:"RDS_DATABASE"`
		Port            int           `env:"RDS_PORT"`
		MaxIdleConns    int           `env:"RDS_MAX_IDLE_CONNS"`
		MaxOpenConns    int           `env:"RDS_MAX_OPEN_CONNS"`
		ConnMaxIdleTime time.Duration `env:"RDS_CONN_MAX_IDLE_TIME"`
		ConnMaxLifetime time.Duration `env:"RDS_CONN_MAX_LIFETIME"`
	}
)

func New() (*Config, error) {
	var cfg *Config

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return &Config{}, err
	}

	err = envdecode.StrictDecode(cfg)
	if err != nil {
		return &Config{}, err
	}

	err = cfg.Validate()
	if err != nil {
		return &Config{}, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	// Ensure the appropriate values are set based on the storage
	switch c.App.Storage {
	case StorageDynamoDB:
		if c.AWS.AccessKeyID == "" || c.AWS.Secret == "" || c.AWS.Region == "" {
			return errors.New("missing required AWS configuration")
		} else if c.DynamoDB.ItemTable == "" {
			return errors.New("missing required DynamoDB configuration")
		}
	default:
	}
	return nil
}

func loadAWSCfg(ctx context.Context, cfg Config) (aws.Config, error) {
	customResolver := aws.EndpointResolverWithOptions(
		aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if cfg.AWS.Endpoint != "" {
					return aws.Endpoint{
						URL:           cfg.AWS.Endpoint,
						SigningRegion: region,
						Source:        aws.EndpointSourceCustom,
					}, nil
				}
				// returning EndpointNotFoundError will allow the service to fallback to its default resolution
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		),
	)
	return awsconfig.LoadDefaultConfig(ctx, awsconfig.WithEndpointResolverWithOptions(customResolver))
}
