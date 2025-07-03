package module

import (
	"github.com/finlleylchat/auth/internal/config"
)

func NewConfig() (*config.Config, error) {
	return config.InitConfig()
}
