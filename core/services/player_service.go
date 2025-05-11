package services

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
)

type PlayerService interface {
	AddPlayer(ctx context.Context, token string) error
	GetPlayerByID(ctx context.Context, playerID string) (*models.Player, error)
}

type PlayerServiceImpl struct {
	playerRepo repositories.PlayerRepository
	userRepo   repositories.UserRepository
}

func NewPlayerService(playerRepo repositories.PlayerRepository, userRepo repositories.UserRepository) PlayerService {
	return &PlayerServiceImpl{
		playerRepo: playerRepo,
		userRepo:   userRepo,
	}
}

func (p *PlayerServiceImpl) AddPlayer(ctx context.Context, token string) error {
	if token == "" {
		return exceptions.ErrTokenIsEmpty
	}
	// Extract payload from the token
	payload, err := p.userRepo.ExactJWT(token)
	if err != nil {
		return err
	}

	user, err := p.userRepo.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		return exceptions.ErrUserNotFound
	}
	if user.Nickname == nil {
		return exceptions.ErrNicknameNotFound
	}
	secretCodeNickname := user.Nickname
	secretCodeNumber := fmt.Sprintf("%04d", rand.Intn(10000))

	// Generate a new UUID for the player
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	var player = models.Player{
		ID:         uuid.String(),
		SecretCode: *secretCodeNickname + secretCodeNumber,
		VictimID:   nil,
		IsAlive:    true,
		Score:      0,
		ActionID:   nil,
		Email:      payload.Email,
	}

	if err := p.playerRepo.AddPlayer(ctx, &player); err != nil {
		return err
	}

	return nil
}

func (p *PlayerServiceImpl) GetPlayerByID(ctx context.Context, playerID string) (*models.Player, error) {
	if playerID == "" {
		return nil, exceptions.ErrPlayerIDIsEmpty
	}

	player, err := p.playerRepo.GetPlayerByID(ctx, playerID)
	if err != nil {
		return nil, err
	}

	return player, nil
}
