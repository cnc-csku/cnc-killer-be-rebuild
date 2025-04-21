package services

import (
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
	"github.com/gofiber/contrib/websocket"
)

type ManagerService interface {
	AddPlayer(playerID string, conn *websocket.Conn) (*responses.Message, error)
	RemovePlayer(playerID string)
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
	player, err := m.repo.AddPlayer(playerID, conn)
	if err != nil {
		return nil, err
	}

	log.Printf("player id : %s has added", player.ID)

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
