// Package config - config
package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct for config gRPC-service.
type Config struct {
	Host string `env:"SERVICE_HOST" env-default:"localhost"`
	Port string `env:"SERVICE_PORT" env-default:"8080"`
}

// GetConfig - get env vars.
func GetConfig() *Config {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Println(err)
		help, err := cleanenv.GetDescription(cfg, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println(help)
	}

	return cfg
}
