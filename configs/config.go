package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Env             string `default:"dev"       envconfig:"ENV"`
	Port            string `default:"8081"      envconfig:"PORT"`
	MongoUri        string `required:"true" envconfig:"MONGO_URI"`
	DBName          string `default:"vehiclesapp" envconfig:"DB_NAME"`
	CollectionName  string `default:"vehicles" envconfig:"COLLECTION_NAME"`
	RedisConnection string `default:"localhost:6379" envconfig:"REDIS_CONNECTION"`
}

func Load(logger zerolog.Logger) (*Config, error) {
	var cfg Config

	// Load the .env file from the current working directory
	err := godotenv.Load()
	if err != nil {
		logger.Error().Msgf("Error loading .env file: %s", err)
		return nil, err
	}

	err = envconfig.Process("APP", &cfg)
	if err != nil {
		logger.Error().Msgf("Error while loading the environment config - %s", err)

		return nil, err
	}

	log.Printf("%v", cfg)
	return &cfg, err
}
