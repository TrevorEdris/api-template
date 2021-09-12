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
)

type App struct {
	log    log15.Logger
	server *http.Server
}

func New(cfg ServerConfig, log log15.Logger) *App {
	log.Info("Initializing application", "serverConfig", cfg)
	return &App{
		log:    log,
		server: &http.Server{
			Addr: fmt.Sprintf("0.0.0.0:%d", cfg.Port),
			ReadTimeout: time.Second*time.Duration(cfg.ReadTimeout),
			WriteTimeout: time.Second*time.Duration(cfg.WriteTimeout),
			IdleTimeout: time.Second*time.Duration(cfg.IdleTimeout),
			Handler: api.NewRouter(cfg.JwtIssuer, log),
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
