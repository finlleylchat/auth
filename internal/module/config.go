package module

import (
	"go.uber.org/zap"

	"github.com/finlleylchat/auth/internal/config"
)

func NewConfig(logger *zap.SugaredLogger) *config.Config {
	cfg, err := config.InitConfig(logger)
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}
	logger.Info("Configuration loaded successfully")
	return cfg
}
