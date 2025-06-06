package services

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/google/uuid"
)

type ActionService interface {
	AddAction(ctx context.Context, req *requests.AddActionRequest) error
	FindActionByID(ctx context.Context, actionID string) (*models.Action, error)
}

type actionServiceImpl struct {
	repo repositories.ActionRepository
}

func NewActionService(repo repositories.ActionRepository) ActionService {
	return &actionServiceImpl{
		repo: repo,
	}
}

// Implementation of every methods in `UserService`
func (a *actionServiceImpl) AddAction(ctx context.Context, req *requests.AddActionRequest) error {
	if req == nil {
		return exceptions.ErrInvalidAction
	}

	// Generate a new UUID Version 7 for action_id
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	var action = models.Action{
		ID:        uuid.String(),
		Detail:    req.Detail,
		Condition: req.Condition,
	}

	if err := a.repo.AddAction(ctx, &action); err != nil {
		return err
	}

	return nil
}

func (a *actionServiceImpl) FindActionByID(ctx context.Context, actionID string) (*models.Action, error) {
	if actionID == "" {
		return nil, exceptions.ErrEmptyActionID
	}

	action, err := a.repo.FindActionByID(ctx, actionID)
	if err != nil {
		return nil, err
	}

	return action, nil
}
