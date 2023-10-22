package app

import (
	"os"
	"os/signal"
	"syscall"

	"effective/config"
	v1 "effective/internal/controller/http/v1"
	"effective/internal/repository/postgres"
	"effective/internal/usecase"
	"effective/pkg/httpserver"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Run starts application.
func Run(cfg *config.Config) {
	log.Info().Msg("starting server")

	// Postgres
	db, err := postgres.Connect(cfg.Postgres.DSN, &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("error while connecting to PostgreSQL")
	}
	log.Info().Msg("postgresql connected")

	// Migrate
	err = postgres.Migrate(db)
	if err != nil {
		log.Fatal().Err(err).Msg("error while migrating PostgreSQL")
	} else {
		log.Info().Msg("postgresql migrated")
	}

	// Repository
	humansRepo := postgres.NewHumansRepository(db)

	// UseCase
	humansUC := usecase.NewHumansUseCase(humansRepo)

	// Router
	r := v1.NewRouter(humansUC)

	// HTTP server
	log.Info().Msg("running...")
	httpServer := httpserver.New(r, httpserver.Port(cfg.HTTP.Port))

	// Graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msgf("received stop signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Error().Err(err).Msg("received error from http server")
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Fatal().Err(err).Msg("could not shutdown the server")
	}
}
