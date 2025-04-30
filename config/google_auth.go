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
	if cfg.Google.ClientID == "" || cfg.Google.ClientSecret == "" {
		log.Fatal("error google client ID or google client secret not foud ❌⛔️")
	}

	if cfg.Google.RedirectURL == "" {
		log.Fatal("Redirect URL not found")
	}

	return &GoogleAuthConfig{
		AuthConfig: &oauth2.Config{
			RedirectURL:  cfg.Google.RedirectURL,
			ClientID:     cfg.Google.ClientID,
			ClientSecret: cfg.Google.ClientSecret,
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint: google.Endpoint,
		},

		Config: cfg,
		States: make(map[string]bool),
	}
}
