package manager

import (
	"encoding/json"
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/gofiber/contrib/websocket"
)

func NewGame() *Game {
	return &Game{
		Status:    GameStatusWaiting,
		Players:   make(map[string](*Player)),
		Broadcast: make(chan Message),
	}
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
