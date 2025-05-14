package services

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/gofiber/contrib/websocket"
)

type ManagerService interface {
	AddPlayer(playerID string, conn *websocket.Conn) (*models.Message, error)
	KillPlayer(killerID string, victimID string) error
	RemovePlayer(playerID string)
	ChangeGameStatus(newStatus string) error
	HandleBoardcast() error
	HandlePlayerMessage(playerID string, msgBytes []byte) error
}

type managerServiceImpl struct {
	playerRepo  repositories.PlayerRepository
	managerRepo repositories.ManagerRepository
}

func NewManagerService(repo repositories.ManagerRepository) ManagerService {
	return &managerServiceImpl{
		managerRepo: repo,
	}
}

// AddPlayer implements ManagerService.
func (m *managerServiceImpl) AddPlayer(playerID string, conn *websocket.Conn) (*models.Message, error) {
	m.managerRepo.AddPlayer(playerID, conn)

	gameStatus := m.managerRepo.GetGameStatus()

	return &models.Message{
		Type: requests.MsgTypeUpdateStatus,
		Messages: models.JSON{
			"status": gameStatus,
		},
	}, nil
}

// RemovePlayer implements ManagerService.
func (m *managerServiceImpl) RemovePlayer(playerID string) {
	m.managerRepo.RemovePlayer(playerID)
}

// ChangeGameStatus implements ManagerService.
func (m *managerServiceImpl) ChangeGameStatus(newStatus string) error {
	if _, ok := requests.ValidGameStatus[newStatus]; !ok {
		return exceptions.ErrInvalidGameStatus
	}

	m.managerRepo.ChangeGameStatus(newStatus)
	return nil
}

// HandleBoardcast implements ManagerService.
func (m *managerServiceImpl) HandleBoardcast() error {
	return m.managerRepo.Broadcast()
}

// HandlePlayerMessage implements ManagerService.
func (m *managerServiceImpl) HandlePlayerMessage(playerID string, msgBytes []byte) error {
	return m.managerRepo.PlayerMessageHandle(playerID, msgBytes)
}

// KillPlayer implements ManagerService.
func (m *managerServiceImpl) KillPlayer(killerID string, victimID string) error {
	ctx := context.Background()
	killer, err := m.playerRepo.GetPlayerByID(ctx, killerID)
	if err != nil {
		return err
	}
	victim, err := m.playerRepo.GetPlayerByID(ctx, victimID)
	if err != nil {
		return err
	}
	if !victim.IsAlive {
		return exceptions.ErrPlayerAlreadyDead
	}
	if killer.VictimID == nil || *(killer.VictimID) != victimID {
		return exceptions.ErrInvalidVictim
	}
	err = m.playerRepo.UpdatePlayerIsAlive(ctx, victimID, false)
	if err != nil {
		return err
	}
	m.managerRepo.KillPlayer(killerID, victimID)
	return nil
}
