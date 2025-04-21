package models

import (
	"sync"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
	"github.com/gofiber/contrib/websocket"
)

type JSON map[string]interface{}

type Player struct {
	ID   string
	Conn *websocket.Conn
}

type Game struct {
	Status    string
	Players   map[string](*Player)
	GameMux   sync.RWMutex
	Broadcast chan responses.Message
}
