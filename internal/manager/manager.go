package manager

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type JsonMap map[string]string

const (
	GameStatusWaiting = "waiting"
	GameStatusStart   = "start"
	GameStatusEnd     = "end"
)

var ValidGameStatus = map[string]bool{
	GameStatusWaiting: true,
	GameStatusStart:   true,
	GameStatusEnd:     true,
}

const (
	MsgTypeUpdateStatus = "update-game-status"
	MsgTypeAction       = "actions"
	MsgTypeKill         = "kill"
	MsgTypeRevive       = "revive"
)

type Player struct {
	ID   string
	Conn *websocket.Conn
}

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

type Game struct {
	Status    string
	Players   map[string](*Player)
	GameMux   sync.RWMutex
	Broadcast chan Message
}
