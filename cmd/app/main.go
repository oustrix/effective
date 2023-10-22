package main

import (
	"effective/config"
	"effective/internal/app"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("error while creating app config")
	}

	app.Run(cfg)
}
