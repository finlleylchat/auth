package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"go.uber.org/fx"

	"github.com/finlleylchat/auth/internal/module"
)

func main() {
	app := fx.New(
		fx.Provide(
			module.NewLogger,
			module.NewConfig,
			module.NewDB,
		),
		fx.Invoke(module.StartServer),
	)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		fmt.Println(color.YellowString("\nReceived shutdown signal, shutting down gracefully..."))

		// Даем 30 секунд на graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := app.Stop(ctx); err != nil {
			fmt.Printf(color.RedString("Error during shutdown: %v\n"), err)
			os.Exit(1)
		}

		fmt.Println(color.GreenString("Shutdown completed successfully"))
	}()

	// Запускаем приложение
	if err := app.Err(); err != nil {
		fmt.Printf(color.RedString("Failed to start application: %v\n"), err)
		os.Exit(1)
	}

	app.Run()
}
