package config

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleConfig(cfg Config) oauth2.Config {
	if cfg.GoogleClientID == "" || cfg.GoogleClientSecret == "" {
		log.Fatal("error google client ID or google client secret not foud ❌⛔️")
	}

	if cfg.RedirectURL == "" {
		log.Fatal("Redirect URL not found")
	}

	config := oauth2.Config{
		RedirectURL:  cfg.RedirectURL,
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return config
}
