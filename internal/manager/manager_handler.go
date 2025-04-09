package manager

import (
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type GameHandler interface {
	ChangeGameStatus(c *fiber.Ctx) error
	SubscribePlater(c *websocket.Conn)
}

type gameHandler struct {
	service GameService
}

func NewGameHandler(service GameService) GameHandler {
	return &gameHandler{
		service: service,
	}
}

// ChangeGameStatus implements GameHandler.
func (g *gameHandler) ChangeGameStatus(c *fiber.Ctx) error {
	var body requests.GameStatusRequest
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := g.service.ChangeGameStatus(body.Status)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": g.service.GetGameStatus()})
}

// SubscribePlater implements GameHandler.
func (g *gameHandler) SubscribePlater(c *websocket.Conn) {
	playerID := c.Params("playerID")

	err := g.service.AddPlayer(playerID, c)
	if err != nil {
		g.service.RemovePlayer(playerID)
		c.WriteJSON(fiber.Map{
			"message": "error while adding player",
		})
	}

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from player %s: %v", playerID, err)
			g.service.RemovePlayer(playerID)
			break
		}

		if messageType == websocket.TextMessage {
			g.service.HandlePlayerMessage(playerID, message)
		}
	}

	go g.service.HandleBoardcast()
}
