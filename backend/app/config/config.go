package config

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	HTTP struct {
		Port string `envconfig:"HTTP_PORT" required:"true"`
	}
	Database struct {
		DSN string `envconfig:"DB_DSN" required:"true"`
	}
}

// NewConfig	Считываем настройки в структуру из переменных окружения
func NewConfig() (*AppConfig, error) {
	cfg := new(AppConfig)
	if readError := envconfig.Process("", cfg); readError != nil {
		return cfg, readError
	}
	return cfg, nil
}
