package services

import (
	"context"
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	GetAuthURL() (string, error)
	GetUserInfo(c *fiber.Ctx) (*responses.GoogleResponse, error)
	GetRefreshToken(ctx context.Context, req *requests.TokenRequest) (string, error)
}

type authServiceImpl struct {
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
}

func NewAuthService(authRepo repositories.AuthRepository, userRepo repositories.UserRepository) AuthService {
	return &authServiceImpl{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

// GetAuthURl implements AuthService.
func (a *authServiceImpl) GetAuthURL() (string, error) {
	state, err := a.authRepo.GenerateState()
	if err != nil {
		return "", err
	}
	return a.authRepo.GetAuthURL(state), nil
}

// GetUserInfo implements AuthService.
func (a *authServiceImpl) GetUserInfo(c *fiber.Ctx) (*responses.GoogleResponse, error) {
	ctx := c.Context()
	state := c.Query("state")
	redirect := c.Query("redirect_url") // search param have to be send before api calling
	log.Printf("%s", redirect)
	if state == "" {
		return nil, exceptions.ErrNoState
	}

	if validated := a.authRepo.VerifyState(state); !validated {
		return nil, exceptions.ErrInvalidState
	}

	code := c.Query("code")
	if code == "" {
		return nil, exceptions.ErrCodeNotFound
	}

	token, err := a.authRepo.ExchangeCode(ctx, code)

	if err != nil {
		return nil, exceptions.ErrExchangeFailed
	}

	googleUser, err := a.authRepo.GetUserInfo(ctx, token)

	if err != nil {
		return nil, exceptions.ErrFetchGoogleUser
	}

	user, err := a.userRepo.FindUserByEmail(ctx, googleUser.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = a.userRepo.AddUser(ctx, googleUser.Email)
		if err != nil {
			return nil, err
		}
	}

	signedToken, err := a.userRepo.GenerateAccessToken(user, true)
	if err != nil {
		return nil, err
	}

	return &responses.GoogleResponse{
		Name:       googleUser.Name,
		Email:      googleUser.Email,
		PictureURL: googleUser.PictureURL,
		Token:      signedToken,
	}, nil

}

// GetRefreshToken implements AuthService.
func (a *authServiceImpl) GetRefreshToken(ctx context.Context, req *requests.TokenRequest) (string, error) {
return a.userRepo.GenerateRefreshToken(ctx, req.Token)
}
