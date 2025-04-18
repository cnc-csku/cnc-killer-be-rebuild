package services

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
)

type GoogleAuthInstance struct {
	GoogleAuth *config.GoogleAuthConfig
}

func NewGoogleAuthInstance(cfg *config.GoogleAuthConfig) repositories.AuthRepository {
	return &GoogleAuthInstance{
		GoogleAuth: cfg,
	}
}

// GenerateState implements repositories.AuthRepository.
func (g *GoogleAuthInstance) GenerateState() (string, error) {
	state, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return state.String(), err
}

// GetAuthURL implements repositories.AuthRepository.
func (g *GoogleAuthInstance) GetAuthURL() error {
	state, err := g.GenerateState()
	if err != nil {
		return err
	}
}

// VerifyState implements repositories.AuthRepository.
func (g *GoogleAuthInstance) VerifyState(state string) bool {
	if _, ok := g.GoogleAuth.States[state]; ok {
		delete(g.GoogleAuth.States, state)
		return true
	} else {
		return false
	}
}
