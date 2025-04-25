package services

import (
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	GetAuthURL() (string, error)
	GetUserInfo(c *fiber.Ctx) (*responses.GoogleResponse, error)
}

type authServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authServiceImpl{
		repo: repo,
	}
}

// GetAuthURl implements AuthService.
func (a *authServiceImpl) GetAuthURL() (string, error) {
	state, err := a.repo.GenerateState()
	if err != nil {
		return "", err
	}
	return a.repo.GetAuthURL(state), nil
}

// GetUserInfo implements AuthService.
func (a *authServiceImpl) GetUserInfo(c *fiber.Ctx) (*responses.GoogleResponse, error) {
	ctx := c.Context()
	state := c.Query("state")
	redirect := c.Query("redirect_url")
	log.Printf("%s", redirect)
	if state == "" {
		return nil, exceptions.ErrNoState
	}

	if validated := a.repo.VerifyState(state); !validated {
		return nil, exceptions.ErrInvalidState
	}

	code := c.Query("code")
	if code == "" {
		return nil, exceptions.ErrCodeNotFound
	}

	token, err := a.repo.ExchangeCode(ctx, code)

	if err != nil {
		return nil, exceptions.ErrExchangeFailed
	}

	user, err := a.repo.GetUserInfo(ctx, token)

	if err != nil {
		return nil, exceptions.ErrFetchGoogleUser
	}

	return &responses.GoogleResponse{
		Name:       user.Name,
		Email:      user.Email,
		PictureURL: user.PictureURL,
	}, nil

}
