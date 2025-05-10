package models

import (
	"github.com/gofiber/contrib/websocket"
)

type JSON map[string]interface{}

type Manager struct {
	ID   string
	Conn *websocket.Conn
}

type Message struct {
	Type     string                 `json:"type"`
	Messages map[string]interface{} `json:"messages"`
}
