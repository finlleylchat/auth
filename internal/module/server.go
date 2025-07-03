package module

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/finlleylchat/auth/internal/config"
)

func StartServer(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, cfg *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting auth service",
				zap.String("db_host", cfg.DB.Host),
				zap.Int("db_port", cfg.DB.Port),
			)
			fmt.Println(color.GreenString("Auth service started successfully!"))
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping auth service")
			return nil
		},
	})
}
