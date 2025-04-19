package facilities

import (
	"context"
	"encoding/json"
	"io"
	"log"

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
	g.GoogleAuth.States[state] = true
	return g.GoogleAuth.AuthConfig.AuthCodeURL(state)
}

// VerifyState implements repositories.AuthRepository.
func (g *GoogleAuthInstance) VerifyState(state string) bool {
	if g.GoogleAuth.States[state] {
		delete(g.GoogleAuth.States, state)
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

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var googleUser models.Google
	userData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(userData, &googleUser)
	log.Printf("err : %s", err)
	if err != nil {
		return nil, err
	}

	return &googleUser, nil

}
