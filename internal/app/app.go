// Package app configures and runs application.
package app

import (
	"github.com/Nikola-zim/3d-printing-studio/config"
	v0 "github.com/Nikola-zim/3d-printing-studio/internal/controller/http/v0"
	"github.com/Nikola-zim/3d-printing-studio/internal/usecase"
	"github.com/Nikola-zim/3d-printing-studio/pkg/httpserver"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

const (
	ReleaseModTrue = 1
)

// Run creates objects via constructors.
func Run(cfg config.Config, log zerolog.Logger) {
	// HTTP Server
	if cfg.ReleaseMod == ReleaseModTrue {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	//ordersManager usecase
	orderManager := usecase.NewOrderManager()

	handler := gin.New()
	v0.NewRouter(handler, log, *orderManager)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
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
