package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT"`
}

func NewConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env file not found 🚨")
	}
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		log.Fatal("error while try to parse .env file")
	}

	return cfg
}
