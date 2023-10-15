package app

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/TrevorEdris/api-template/app/config"
	"github.com/TrevorEdris/api-template/app/controller"
	"github.com/TrevorEdris/api-template/app/infrastructure"
	"github.com/TrevorEdris/api-template/app/repository"
	"github.com/TrevorEdris/api-template/app/service"
)

var (
	c             *container
	containerOnce sync.Once
)

type (
	Container interface {
		Config() *config.Config
		Logger() *zap.Logger
		ItemController() controller.ItemController
	}

	container struct {
		cfg    *config.Config
		logger *zap.Logger
	}
)

func NewContainer() Container {
	if c == nil {
		containerOnce.Do(func() {
			c = &container{}
			err := c.init()
			if err != nil {
				panic(err)
			}
		})
	}
	return c
}

func (c *container) init() error {
	err := c.initConfig()
	if err != nil {
		return err
	}

	err = c.initLogging()
	if err != nil {
		return err
	}

	return nil
}

func (c *container) initConfig() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	c.cfg = cfg
	return nil
}

func (c *container) initLogging() error {
	prodCfg := zap.NewProductionConfig()
	atomicLevel := zap.NewAtomicLevel()
	err := atomicLevel.UnmarshalText([]byte(c.cfg.LogLevel))
	if err != nil {
		return fmt.Errorf("failed to parse log level '%s': %w", c.cfg.LogLevel, err)
	}
	prodCfg.Level = atomicLevel

	logger, err := prodCfg.Build()
	if err != nil {
		return fmt.Errorf("failed to build logger from config: %w", err)
	}
	c.logger = logger
	return nil
}

func (c *container) Config() *config.Config {
	return c.cfg
}

func (c *container) Logger() *zap.Logger {
	return c.logger
}

func (c *container) ItemController() controller.ItemController {
	pDriver := infrastructure.NewPostgresDriver(&c.cfg.RDS)
	itemRepo := repository.NewItemRepository(pDriver)
	itemService := service.NewItemService(itemRepo)
	itemController := controller.NewItemController(itemService)

	return itemController
}
