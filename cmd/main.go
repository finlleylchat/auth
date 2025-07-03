package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"go.uber.org/fx"

	"github.com/finlleylchat/auth/internal/module"
)

func main() {
	app := fx.New(
		fx.Provide(
			module.NewLogger,
			module.NewConfig,
		),
		fx.Invoke(module.StartServer),
	)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		fmt.Println(color.YellowString("\nShutting down gracefully..."))
		app.Stop(context.Background())
	}()

	app.Run()
}
