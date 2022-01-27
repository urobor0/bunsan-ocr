package bootstrap

import (
	"bunsan-ocr/internal/platform/bus/inmemory"
	"bunsan-ocr/internal/platform/server"
	"context"
	"github.com/kelseyhightower/envconfig"
	"time"
)

func Run() error {
	var cfg config
	err := envconfig.Process("DDD", &cfg)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus = inmemory.NewQueryBus()
	)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"root"`
	DbPass    string        `default:"123456"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"5432"`
	DbName    string        `default:"ocr"`
	DbTimeout time.Duration `default:"5s"`
}
