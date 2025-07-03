package database

import (
	"context"
	"time"

	"github.com/finlleylchat/auth/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func InitDB(ctx context.Context, config *config.Config, logger *zap.SugaredLogger) (*sqlx.DB, error) {
	logger.Info("Connecting to database")

	dsn := config.DB.DSN()

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Настраиваем connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// Проверяем подключение с timeout
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		logger.Errorf("Failed to ping database: %v", err)
		db.Close()
		return nil, err
	}

	logger.Info("Successfully connected to database")
	return db, nil
}
