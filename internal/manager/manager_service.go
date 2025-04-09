package manager

import (
	"encoding/json"
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/gofiber/contrib/websocket"
)

type GameService interface {
	AddPlayer(userID string, conn *websocket.Conn) error
	GetGameStatus() string
	RemovePlayer(userID string)
	ChangeGameStatus(newStatus string) error
	HandleBoardcast() error
	HandlePlayerMessage(userID string, msgBytes []byte) error
}

func NewGame() GameService {
	return &Game{
		Status:    GameStatusWaiting,
		Players:   make(map[string](*Player)),
		Broadcast: make(chan Message),
	}
}

func (g *Game) GetGameStatus() string {
	return g.Status
}

// AddPlayer implements GameService.
func (g *Game) AddPlayer(userID string, conn *websocket.Conn) error {
	player := &Player{
		ID:   userID,
		Conn: conn,
	}

	g.GameMux.Lock()
	g.Players[userID] = player
	g.GameMux.Unlock()

	log.Printf("Player %s joined the game. Total players: %d", userID, len(g.Players))

	statusMsg := Message{
		Type:    "status",
		Content: JsonMap{"status": g.Status},
	}

	msgBytes, err := json.Marshal(statusMsg)
	if err != nil {
		return err
	}

	conn.WriteMessage(websocket.TextMessage, msgBytes)
	return nil
}

func (g *Game) RemovePlayer(userID string) {
	g.GameMux.Lock()
	delete(g.Players, userID)
	g.GameMux.Unlock()

	log.Printf("Player %s left the game. Remaining players: %d", userID, len(g.Players))
}

func (g *Game) ChangeGameStatus(newStatus string) error {
	_, ok := ValidGameStatus[newStatus]
	if !ok {
		return exceptions.ErrInvalidGameStatus
	}
	g.Status = newStatus

	log.Printf("Game status changed to: %s", newStatus)

	statusMsg := Message{
		Type: MsgTypeUpdateStatus,
		Content: JsonMap{
			"status": newStatus,
		},
	}

	g.Broadcast <- statusMsg
	return nil

}

func (g *Game) HandleBoardcast() error {
	for {
		msg := <-g.Broadcast

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			return exceptions.ErrConvertJSON
		}

		g.GameMux.RLock()

		for _, player := range g.Players {
			err := player.Conn.WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
				log.Printf("Error sending message to player %s: %v", player.ID, err)
				return err
			}
		}

		g.GameMux.RUnlock()
		log.Printf("Broadcasted message of type '%s' to %d players", msg.Type, len(g.Players))
	}
}

// HandlerPlayerMessage implements GameService.
func (g *Game) HandlePlayerMessage(userID string, msgBytes []byte) error {
	var msg map[string]interface{}
	if err := json.Unmarshal(msgBytes, &msg); err != nil {
		return err
	}
	msgType, ok := msg["type"].(string)
	if !ok {
		return exceptions.ErrInvalidType
	}
	switch msgType {
	case MsgTypeUpdateStatus:
		newStatus, ok := msg["content"].(string)
		if !ok {
			return exceptions.ErrInvalidRequest
		}
		g.ChangeGameStatus(newStatus)
	case MsgTypeKill:
		return nil
	case MsgTypeRevive:
		return nil
	default:
		return exceptions.ErrInvalidType
	}

	return exceptions.ErrInvalidRequest
}
