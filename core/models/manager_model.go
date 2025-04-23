package models

import (
	"github.com/gofiber/contrib/websocket"
)

type JSON map[string]interface{}

type Player struct {
	ID   string
	Conn *websocket.Conn
}
