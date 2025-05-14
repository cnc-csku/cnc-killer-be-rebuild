package repositories

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/gofiber/contrib/websocket"
)

type ManagerRepository interface {
	AddPlayer(playerID string, conn *websocket.Conn) *models.Manager
	KillPlayer(killerID string, victimID string)
	RemovePlayer(playerID string)
	GetGameStatus() string
	ChangeGameStatus(newStatus string)
	Broadcast() error
	PlayerMessageHandle(playerID string, msgBytes []byte) error
}
