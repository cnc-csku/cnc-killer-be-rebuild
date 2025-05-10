package services

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
)

type UserService interface {
	GetUserRole(ctx context.Context, email string) (*responses.RoleResponse, error)
	ChangeUserNickname(ctx context.Context, req *requests.ChangeNicknameRequest) error
}

type userServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

// GetUserRole implements UserService.
func (u *userServiceImpl) GetUserRole(ctx context.Context, email string) (*responses.RoleResponse, error) {
	user, err := u.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, exceptions.ErrUserNotFound
	}

	return &responses.RoleResponse{
		Role: user.Role,
	}, nil
}

// ChangeUserNickname implements UserService.
func (u *userServiceImpl) ChangeUserNickname(ctx context.Context, req *requests.ChangeNicknameRequest) error {
	return u.repo.UpdateUserNickname(ctx, req.Email, req.Nickname)
}
