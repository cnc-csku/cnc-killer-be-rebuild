package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port uint16 `env:"PORT"`
	// for postgres
	DBHost     string `env:"DB_HOST"`
	DBPort     uint16 `env:"DB_PORT"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBDatabase string `env:"DB_DATABASE"`
	DBSSLMode  string `env:"DB_SSL_MODE"`
	// for google
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
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
