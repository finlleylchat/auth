package main

import (
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/fx"

	"github.com/finlleylchat/auth/internal/module"
)

func main() {
	fx.New(
		fx.Provide(module.NewConfig()),
	).Run()
	fmt.Println(color.GreenString("Hello, world!"))
}
