package config

import (
	"sync"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
)

type Game struct {
	Status    string
	Players   map[string](*models.Manager)
	GameMux   sync.RWMutex
	Broadcast chan models.Message
}

func NewGame() *Game {
	return &Game{
		Status:    requests.GameStatusWaiting,
		Players:   make(map[string](*models.Manager)),
		Broadcast: make(chan models.Message),
	}
}
