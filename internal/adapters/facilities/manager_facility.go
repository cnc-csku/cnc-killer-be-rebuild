package facilities

import (
	"encoding/json"
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ManagerInstance struct {
	game *config.Game
}

func NewManagerInstance(game *config.Game) repositories.ManagerRepository {
	return &ManagerInstance{
		game: game,
	}
}

// AddPlayer implements repositories.ManagerRepository.
func (m *ManagerInstance) AddPlayer(playerID string, conn *websocket.Conn) *models.Manager {
	player := &models.Manager{
		ID:   playerID,
		Conn: conn,
	}

	m.game.GameMux.Lock()
	m.game.Players[playerID] = player
	m.game.GameMux.Unlock()

	log.Printf("Player %s joined the game. Total players: %d", playerID, len(m.game.Players))

	return player
}

// Broadcast implements repositories.ManagerRepository.
func (m *ManagerInstance) Broadcast() error {
	for {
		msg := <-m.game.Broadcast
		log.Print(msg)

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			return exceptions.ErrConvertJSON
		}

		m.game.GameMux.RLock()

		for _, player := range m.game.Players {
			err := player.Conn.WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
				log.Printf("Error sending message to player %s: %v", player.ID, err)
				return err
			}
		}

		m.game.GameMux.RUnlock()
		log.Printf("Broadcasted message of type '%s' to %d players", msg.Type, len(m.game.Players))
	}
}

// ChangeGameStatus implements repositories.ManagerRepository.
func (m *ManagerInstance) ChangeGameStatus(newStatus string) {
	m.game.Status = newStatus
	statusMsg := models.Message{
		Type:     requests.MsgTypeUpdateStatus,
		Messages: models.JSON{"status": newStatus},
	}

	m.game.Broadcast <- statusMsg
}

// GetGameStatus implements repositories.ManagerRepository.
func (m *ManagerInstance) GetGameStatus() string {
	return m.game.Status
}

// PlayerMessageHandle implements repositories.ManagerRepository.
func (m *ManagerInstance) PlayerMessageHandle(playerID string, msgBytes []byte) error {
	var msg models.Message
	if err := json.Unmarshal(msgBytes, &msg); err != nil {
		return err
	}
	switch msg.Type {
	case requests.MsgTypeKill:
	case requests.MsgTypeRevive:
		return nil
	default:
		return exceptions.ErrInvalidType
	}

	return exceptions.ErrInvalidRequest
}

// RemovePlayer implements repositories.ManagerRepository.
func (m *ManagerInstance) RemovePlayer(playerID string) {
	m.game.GameMux.Lock()
	delete(m.game.Players, playerID)
	m.game.GameMux.Unlock()

	log.Printf("Player %s left the game. Remaining players: %d", playerID, len(m.game.Players))
}

// KillPlayer implements repositories.ManagerRepository.
func (m *ManagerInstance) KillPlayer(killerID string, victimID string) {
	victim, ok := m.game.Players[victimID]
	if !ok {
		return
	}
	victim.Conn.WriteJSON(&models.Message{
		Type: requests.MsgTypeKill,
		Messages: fiber.Map{
			"killed_by": killerID,
		},
	})
}
