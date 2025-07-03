package config

import "fmt"

type Config struct {
	DB DBConfig `koanf:"db"`
}

type DBConfig struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Name     string `koanf:"name"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	SSLMode  string `koanf:"sslmode"`
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Name, c.User, c.Password, c.SSLMode)
}

func (c *DBConfig) IsValid() bool {
	return c.Host != "" && c.Port > 0 && c.Name != "" && c.User != "" && c.Password != ""
}
