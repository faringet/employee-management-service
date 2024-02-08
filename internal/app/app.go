package app

import (
	"context"
	"fmt"

	"github.com/engagerocketco/go-common/config"
	"github.com/engagerocketco/go-common/logger"
)

type App struct {
	baseApp *BaseApp
}

func NewApp() (*App, error) {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("main: could not load config: %w", err)
	}

	logger, err := logger.NewLogger(cfg.ServerLogLevel)
	if err != nil {
		return nil, fmt.Errorf("main: could not initialize a logger: %w", err)
	}

	base, err := NewBaseApp(ctx, logger, cfg)
	if err != nil {
		return nil, fmt.Errorf("main: could not init base app: %w", err)
	}

	return &App{baseApp: base}, nil
}

func (a *App) Run() {
	a.baseApp.Run(context.Background())
}
