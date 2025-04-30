package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port   uint16       `env:"PORT"`
	DB     DBConfig     `envPrefix:"DB_"`
	Google GoogleConfig `envPrefix:"GOOGLE_"`
	JWT    JWTConfig    `envPrefix:"JWT_"`
}
type DBConfig struct {
	Host     string `env:"HOST"`
	Port     uint16 `env:"PORT"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
	SSLMode  string `env:"SSL_MODE"`
}

type GoogleConfig struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	UserInfoURL  string `env:"USER_INFO_URL"`
	RedirectURL  string `env:"REDIRECT_URL"`
}

type JWTConfig struct {
	Secret     string `env:"SECRET"`
	AccessExp  string `env:"ACCESS_EXP"`
	RefreshExp string `env:"REFRESH_EXP"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env file not found 🚨")
	}
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		log.Fatal("error while try to parse .env file")
	}

	return &cfg
}
