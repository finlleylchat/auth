package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

func InitConfig(logger *zap.SugaredLogger) (*Config, error) {
	logger.Info("Loading configuration from environment variables")

	if err := godotenv.Load(); err != nil {
		logger.Debug("No .env file found, using system environment variables")
	} else {
		logger.Info(".env file loaded successfully")
	}

	k := koanf.New(".")

	if err := k.Load(env.Provider(".", env.Opt{
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ToLower(k)
			k = strings.ReplaceAll(k, "_", ".")
			return k, v
		},
	}), nil); err != nil {
		return nil, fmt.Errorf("failed to load env config: %w", err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	setDefaults(&cfg)
	logger.Debug("Applied default values",
		zap.String("db_host", cfg.DB.Host),
		zap.Int("db_port", cfg.DB.Port),
		zap.String("db_sslmode", cfg.DB.SSLMode),
	)

	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	logger.Info("Configuration loaded and validated successfully")
	return &cfg, nil
}

func setDefaults(cfg *Config) {
	if cfg.DB.Host == "" {
		cfg.DB.Host = "localhost"
	}
	if cfg.DB.Port == 0 {
		cfg.DB.Port = 5432
	}
	if cfg.DB.SSLMode == "" {
		cfg.DB.SSLMode = "disable"
	}
}

func validateConfig(cfg *Config) error {
	if cfg.DB.Name == "" {
		return fmt.Errorf("database name is required")
	}
	if cfg.DB.User == "" {
		return fmt.Errorf("database user is required")
	}
	if cfg.DB.Password == "" {
		return fmt.Errorf("database password is required")
	}
	return nil
}
