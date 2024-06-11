package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Env            string `default:"dev"       envconfig:"ENV"`
	Port           string `default:"8080"      envconfig:"PORT"`
	MongoUri       string `required:"true" envconfig:"MONGO_URI"`
	DBName         string `default:"vehiclesapp" envconfig:"DB_NAME"`
	CollectionName string `default:"vehicles" envconfig:"COLLECTION_NAME"`
}

func Load(logger zerolog.Logger) (*Config, error) {
	var cfg Config

	err := envconfig.Process("APP", &cfg)
	if err != nil {
		logger.Error().Msgf("Error while loading the environment config - %s", err)

		return nil, err
	}

	log.Printf("%v", cfg)
	return &cfg, err
}