package app

import (
	"context"
	"crypto-rates-client/internal/config"
	"crypto-rates-client/internal/http/delivery"
	"crypto-rates-client/internal/http/repository"
	"time"
)

// Application представляет основное приложение.
type Application struct {
    Config    *config.Config
	Handler    *delivery.Handler
}

// New создает новый объект Application и настраивает его.
func New(ctx context.Context) (*Application, error) {
    // инициализирует конфигурацию
    cfg, err := config.NewConfig(ctx)
    if err != nil {
        return nil, err
    }
    cfg.TTL = time.Minute
    h := delivery.NewHandler(repository.NewRepository(cfg.BaseURL,cfg.TTL))
    // инициализирует объект Application
    app := &Application{
        Config: cfg,
        Handler: h,
    }

    return app, nil
}

