package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/inconshreveable/log15"

	"github.com/TrevorEdris/api-template/app/api"
	"github.com/TrevorEdris/api-template/app/internal/repository"
	"github.com/TrevorEdris/api-template/app/models/item"
	"github.com/TrevorEdris/api-template/somedatabase"
)

const (
	// AppName is the name of this application.
	AppName = "api-template"

	// EnvPrefix is used to narrow the scope of environment variables being parsed to
	// only those that start with this prefix.
	//
	// Example:
	//    1. API_TEMPLATE_SOME_CONFIG_VALUE=true
	//    2. SOME_OTHER_CONFIG_VALUE=false
	//
	// Example 1 _would_ be parsed by this application, as it begins with the below `EnvPrefix`.
	// Example 2 _would not_ be parsed, as it does not begin with the `EnvPrefix`.
	EnvPrefix = "API_TEMPLATE_"
)

// App is a wrapper around the server application.
type App struct {
	log    log15.Logger
	server *http.Server
}

// New creates an instance of an App.
func New(cfg ServerConfig, log log15.Logger, someDBClient *somedatabase.SomeDatabaseClient) *App {
	log.Info("Initializing application", "serverConfig", cfg)

	// Create the model-database interactor
	items := item.New(repository.NewItemSomeDatabase(someDBClient))
	models := api.Models{
		Items: items,
	}

	return &App{
		log: log,
		server: &http.Server{
			Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
			ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
			WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
			IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
			Handler:      api.NewRouter(cfg.JwtIssuer, log, models),
		},
	}
}

// Run creates and starts the HTTP server.
func (a *App) Run() {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// An interrupt signal has been received, so the server will be shutdown
		if err := a.server.Shutdown(context.Background()); err != nil {
			// Some error occured from closing listeners or the context timed out
			a.log.Error("Error shutting down the server", "error", err)
			os.Exit(1)
		}
		close(idleConnsClosed)
	}()

	a.log.Info("Server beginning to listen")
	if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
		a.log.Error("Unable to ListenAndServe", "error", err)
		os.Exit(1)
	}

	<-idleConnsClosed
}
