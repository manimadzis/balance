package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	// Адрес, на котором поднимается сервер
	Address string `toml:"address"`

	// Подключение к базе данных
	DBConnectURL string `toml:"db_connection_string"`
}

func DefaultConfig() *Config {
	return &Config{
		Address:      ":62000",
		DBConnectURL: "postgres://postgres:pass@localhost:5432/accounts?sslmode=disable",
	}
}

func ParseFile(filename string) (*Config, error) {
	config := DefaultConfig()
	if _, err := toml.DecodeFile(filename, config); err != nil {
		return config, err
	}

	return config, nil
}
