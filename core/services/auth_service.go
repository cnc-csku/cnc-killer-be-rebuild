package services

import (
	"context"
	"fmt"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
)

type AuthService interface {
	GetAuthURL() (string, error)
	GetUserInfo(req requests.GoogleAuthInfoRequest, ctx context.Context) (*responses.GoogleResponse, error)
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
	// fmt.Printf("call auth url")
	state, err := a.repo.GenerateState()
	if err != nil {
		return "", err
	}
	return a.repo.GetAuthURL(state), nil
}

// GetUserInfo implements AuthService.
func (a *authServiceImpl) GetUserInfo(req requests.GoogleAuthInfoRequest, ctx context.Context) (*responses.GoogleResponse, error) {
	if req.State == "" {
		return nil, exceptions.ErrNoState
	}

	if validated := a.repo.VerifyState(req.State); !validated {
		return nil, exceptions.ErrInvalidState
	}

	if req.Code == "" {
		return nil, exceptions.ErrCodeNotFound
	}

	token, err := a.repo.ExchangeCode(ctx, req.Code)

	if err != nil {
		return nil, exceptions.ErrExchangeFailed
	}

	user, err := a.repo.GetUserInfo(ctx, token)

	if err != nil {
		fmt.Printf("error : %s", err.Error())
		return nil, exceptions.ErrFetchGoogleUser
	}

	return &responses.GoogleResponse{
		Name:       user.Name + " " + user.FamilyName,
		Email:      user.Email,
		PictureURL: user.PictureURL,
	}, nil

}
