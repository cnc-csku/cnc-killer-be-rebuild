package config

import (
	"sync"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
)

type Game struct {
	Status    string
	Players   map[string](*models.Player)
	GameMux   sync.RWMutex
	Broadcast chan responses.Message
}

func NewGame() *Game {
	return &Game{
		Status:    requests.GameStatusWaiting,
		Players:   make(map[string](*models.Player)),
		Broadcast: make(chan responses.Message),
	}
}
