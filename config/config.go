package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		HTTP
		Postgres
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT" env-required:"true"`
	}

	// Postgres -.
	Postgres struct {
		DSN string `env:"POSTGRES_DSN" env-required:"true"`
	}
)

// New creates new Config instance and returns it.
func New() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
