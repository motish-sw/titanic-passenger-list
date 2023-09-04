package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

//Any env config parameters should be added to the following structure
//A default value should be set here or in a config file
type Configuration struct {
	Address string `env:"ADDRESS" envDefault:":8080"`
}

func NewConfig(files ...string) (*Configuration, error) {
	if err := godotenv.Load(files...); err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
