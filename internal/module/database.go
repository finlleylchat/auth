package module

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/finlleylchat/auth/internal/config"
	"github.com/finlleylchat/auth/internal/database"
)

func NewDB(lc fx.Lifecycle, cfg *config.Config, logger *zap.SugaredLogger) *sqlx.DB {
	var db *sqlx.DB

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			db, err = database.InitDB(ctx, cfg, logger)
			if err != nil {
				logger.Errorf("Failed to initialize database: %v", err)
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if db != nil {
				logger.Info("Closing database connection")
				// Даем время для завершения активных запросов
				shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
				defer cancel()

				done := make(chan error, 1)
				go func() {
					done <- db.Close()
				}()

				select {
				case err := <-done:
					if err != nil {
						logger.Errorf("Error closing database: %v", err)
						return err
					}
					logger.Info("Database connection closed successfully")
				case <-shutdownCtx.Done():
					logger.Warn("Database shutdown timeout exceeded")
					return shutdownCtx.Err()
				}
			}
			return nil
		},
	})

	return db
}
