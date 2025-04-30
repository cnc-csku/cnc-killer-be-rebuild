package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"golang.org/x/oauth2"
)

type AuthRepository interface {
	GetAuthURL(state string) string
	ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*models.GooglePayload, error)
	VerifyState(state string) bool
	GenerateState() (string, error)
}
