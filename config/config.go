package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
	}

	App struct {
		Name     string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version  string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Location string `env-required:"true" yaml:"location" env:"LOCATION"`
	}

	HTTP struct {
		Port       string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		ReleaseMod int    `env-required:"true" yaml:"release_mod" env:"RELEASE_MOD"`
	}

	// PG - Postgres БД.
	PG struct {
		PoolMax      int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL          string `env-required:"true" yaml:"pg_url"   env:"ER_PG_URL"`
		User         string `env-required:"true" yaml:"user"     env:"ER_PG_USER"`
		Password     string `env-required:"true" yaml:"password" env:"ER_PG_PASSWORD"`
		DatabaseName string `env-required:"true" yaml:"db"       env:"ER_PG_DB"`
	}
)

// NewConfig returns app config.
func NewConfig(debug bool) (*Config, error) {
	cfg := &Config{}

	var err error

	if debug {
		err = cleanenv.ReadConfig("./config/config_debug.yml", cfg)
	} else {
		err = cleanenv.ReadConfig("./config/config.yml", cfg)
	}

	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
