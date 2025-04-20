package services

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
	"github.com/google/uuid"
)

type UserService interface {
	GetUserRole(ctx context.Context, userID string) (*responses.RoleResponse, error)
	Login(ctx context.Context, req requests.UserLoginRequest) (*responses.UserResponse, error)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

type userServiceImpl struct {
	repo repositories.UserRepository
}

// GetUserRole implements UserService.
func (u *userServiceImpl) GetUserRole(ctx context.Context, userID string) (*responses.RoleResponse, error) {
	err := uuid.Validate(userID)
	if err != nil {
		return nil, exceptions.ErrInvalidUUID
	}
	user, err := u.repo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &responses.RoleResponse{
		Role: user.Role,
	}, nil
}

// Login implements UserService.
func (u *userServiceImpl) Login(ctx context.Context, req requests.UserLoginRequest) (*responses.UserResponse, error) {
	panic("unimplemented")
}
