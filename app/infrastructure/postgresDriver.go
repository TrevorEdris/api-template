package infrastructure

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	// Importing the postgres driver
	_ "github.com/lib/pq"

	"github.com/TrevorEdris/api-template/app/config"
)

type (
	PostgresDriver interface {
		Ping() error
		SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	}

	postgresDriver struct {
		*sqlx.DB
	}
)

var (
	pd           *postgresDriver
	postgresOnce sync.Once
)

func NewPostgresDriver(cfg *config.RDS) PostgresDriver {
	if pd == nil {
		postgresOnce.Do(func() {
			db, err := sqlx.Connect("postgres", fmt.Sprintf(
				"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
			))
			if err != nil {
				panic(err)
			}

			db.SetMaxIdleConns(cfg.MaxIdleConns)
			db.SetMaxOpenConns(cfg.MaxOpenConns)
			db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
			db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

			pd = &postgresDriver{db}
		})
	}
	return pd
}
