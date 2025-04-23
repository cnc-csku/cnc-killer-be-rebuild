package services

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
	"github.com/gofiber/contrib/websocket"
)

type ManagerService interface {
	AddPlayer(playerID string, conn *websocket.Conn) (*responses.Message, error)
	RemovePlayer(playerID string)
	ChangeGameStatus(newStatus string) error
	HandleBoardcast() error
	HandlePlayerMessage(playerID string, msgBytes []byte) error
}

type managerServiceImpl struct {
	repo repositories.ManagerRepository
}

func NewManagerService(repo repositories.ManagerRepository) ManagerService {
	return &managerServiceImpl{
		repo: repo,
	}
}

// AddPlayer implements ManagerService.
func (m *managerServiceImpl) AddPlayer(playerID string, conn *websocket.Conn) (*responses.Message, error) {
	m.repo.AddPlayer(playerID, conn)

	gameStatus := m.repo.GetGameStatus()

	return &responses.Message{
		MessageType: requests.MsgTypeUpdateStatus,
		Contents: models.JSON{
			"status": gameStatus,
		},
	}, nil
}

// RemovePlayer implements ManagerService.
func (m *managerServiceImpl) RemovePlayer(playerID string) {
	m.repo.RemovePlayer(playerID)
}

// ChangeGameStatus implements ManagerService.
func (m *managerServiceImpl) ChangeGameStatus(newStatus string) error {
	if _, ok := requests.ValidGameStatus[newStatus]; !ok {
		return exceptions.ErrInvalidGameStatus
	}

	m.repo.ChangeGameStatus(newStatus)
	return nil
}

// HandleBoardcast implements ManagerService.
func (m *managerServiceImpl) HandleBoardcast() error {
	return m.repo.Broadcast()
}

// HandlePlayerMessage implements ManagerService.
func (m *managerServiceImpl) HandlePlayerMessage(playerID string, msgBytes []byte) error {
	return m.repo.PlayerMessageHandle(playerID, msgBytes)
}
