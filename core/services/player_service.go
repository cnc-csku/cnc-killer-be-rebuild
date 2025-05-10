package services

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
)

type PlayerService interface {
	AddPlayer(ctx context.Context, token string) error
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
	// Extract JWT token to get payload
	payload, err := p.userRepo.ExactJWT(token)
	if err != nil {
		return err
	}

	// Generate a new UUID for the player
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	secretCodeNickname := "" //TODO: GetNicknameByEmail
	secretCodeNumber := fmt.Sprintf("%04d", rand.Intn(10000))

	var player = models.Player{
		ID:         uuid.String(),
		SecretCode: secretCodeNickname + secretCodeNumber,
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
