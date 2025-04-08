package manager

import (
	"encoding/json"
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/fasthttp/websocket"
)

func (g *Game) HandlerBoardcast() error {
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
func (g *Game) HandlerPlayerMessage(userID string, msgBytes []byte) error {
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
