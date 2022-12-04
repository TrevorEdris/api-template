package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	echopprof "github.com/sevenNt/echo-pprof"
)

type (
	API interface {
		StartAPI()
	}

	api struct{}
)

var (
	a       *api
	apiOnce sync.Once
)

func NewAPI() API {
	if a == nil {
		apiOnce.Do(func() {
			a = &api{}
		})
	}
	return a
}

func (api *api) initAPI() *echo.Echo {
	logger := NewContainer().Logger()

	e := echo.New()
	baseGroup := e.Group("")
	baseGroup.Use(
		echomw.RequestID(),
		echozap.ZapLogger(logger),
	)

	echopprof.Wrap(e)

	itemGroup := baseGroup.Group("/items")
	ic := NewContainer().ItemController()
	itemGroup.GET("/{id}", ic.GetOne)

	return e
}

func (api *api) StartAPI() {
	e := api.initAPI()
	cfg := NewContainer().Config()

	go func() {
		srv := &http.Server{
			Addr:         fmt.Sprintf("%s:%d", cfg.API.Host, cfg.API.Port),
			ReadTimeout:  cfg.API.ReadTimeout,
			WriteTimeout: cfg.API.WriteTimeout,
			Handler:      e,
		}

		e.Logger.Info("Starting server")
		if err := e.StartServer(srv); err != nil {
			e.Logger.Fatalf("Failed to start server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
