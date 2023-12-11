package main

import (
	"flag"
	"github.com/Nikola-zim/3d-printing-studio/config"
	"github.com/Nikola-zim/3d-printing-studio/internal/app"
	"log"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	debug := flag.Bool("debug", false, "switch debug mode")
	flag.Parse()
	logger := zerolog.New(os.Stderr).
		With().Timestamp().Caller().Logger()
	logger.Info().Msgf("Starting http web server")
	// Configuration
	cfg, err := config.NewConfig(*debug)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Run
	app.Run(*cfg, logger)
}
