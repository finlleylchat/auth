package config

import (
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/v2"
)

func InitConfig() (*Config, error) {
	k := koanf.New(".")
	if err := k.Load(env.Provider(".", env.Opt{}), nil); err != nil {
		return nil, err
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
