package config

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuthConfig struct {
	AuthConfig *oauth2.Config
	Config     *Config
	States     map[string]bool
}

func NewGoogleConfig(cfg *Config) *GoogleAuthConfig {
	if cfg.GoogleClientID == "" || cfg.GoogleClientSecret == "" {
		log.Fatal("error google client ID or google client secret not foud ❌⛔️")
	}

	if cfg.RedirectURL == "" {
		log.Fatal("Redirect URL not found")
	}

	return &GoogleAuthConfig{
		AuthConfig: &oauth2.Config{
			RedirectURL:  cfg.RedirectURL,
			ClientID:     cfg.GoogleClientID,
			ClientSecret: cfg.GoogleClientSecret,
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint: google.Endpoint,
		},

		Config: cfg,
		States: make(map[string]bool),
	}
}
