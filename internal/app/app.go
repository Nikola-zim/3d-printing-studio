// Package app configures and runs application.
package app

import (
	"context"
	"github.com/Nikola-zim/3d-printing-studio/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

const (
	ReleaseModTrue  = 1
	ReleaseModFalse = 2
)

// Run creates objects via constructors.
func Run(ctx context.Context, cfg *config.Config, logger zerolog.Logger) {
	// Миграции
	Migrate(cfg.PG.URL, cfg.PG.DatabaseName, cfg.PG.User, cfg.PG.Password)

	// Postgres
	pg, err := postgres.New(cfg.PG.URL, cfg.PG.DatabaseName, cfg.PG.User, cfg.PG.Password, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//
	if err != nil {
		log.Fatal().Err(err).Msg("app - Run - postgres.New")
	}
	defer pg.Close()

	// Use case для работы с БД
	repoPG := repo.New(pg, log)

	exchangeRatesUseCase := usecase.New(
		log,
		repoPG,
	)
	// ExchangeRatesCollector - сборщик данных о курсе валют
	exchangeRatesCollector := collector.NewWorker(
		log,
		repoPG,
		cfg.Location,
		collector.MainSourceURL(cfg.Collector.URL),
		collector.UpdateInterval(cfg.Collector.UpdateInterval),
		collector.UpdateAfterTime(cfg.Collector.UpdateAfterTime),
		collector.HistoryDepth(cfg.Collector.HistoryDepth),
	)

	// HTTP Server
	if cfg.ReleaseMod == ReleaseModTrue {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	handler := gin.New()
	http.NewRouter(handler, log, exchangeRatesUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msg("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Err(err).Msg("app - Run - httpServer.Notify: %w")
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Err(err).Msg("app - Run - httpServer.Shutdown: %w")
	}
}
