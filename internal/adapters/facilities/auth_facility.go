package facilities

import (
	"context"
	"encoding/json"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type GoogleAuthInstance struct {
	GoogleAuth *config.GoogleAuthConfig
}

var userInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"

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
func (g *GoogleAuthInstance) GetAuthURL(state string) string {
	return g.GoogleAuth.AuthConfig.AuthCodeURL(state)
}

// VerifyState implements repositories.AuthRepository.
func (g *GoogleAuthInstance) VerifyState(state string) bool {
	if _, ok := g.GoogleAuth.States[state]; ok {
		// delete(g.GoogleAuth.States, state)
		return true
	} else {
		return false
	}
}

// ExchangeCode implements repositories.AuthRepository.
func (g *GoogleAuthInstance) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return g.GoogleAuth.AuthConfig.Exchange(ctx, code)
}

// GetUserInfo implements repositories.AuthRepository.
func (g *GoogleAuthInstance) GetUserInfo(ctx context.Context, token *oauth2.Token) (*models.Google, error) {
	client := g.GoogleAuth.AuthConfig.Client(ctx, token)
	res, err := client.Get(userInfoURL)
	defel res.Body.Close()

	if err != nil {
		return nil, err
	}
	var googleUser *models.Google
	err = json.NewDecoder(res.Body).Decode(googleUser)
	if err != nil {
		return nil, err
	}

	return googleUser, nil

}
