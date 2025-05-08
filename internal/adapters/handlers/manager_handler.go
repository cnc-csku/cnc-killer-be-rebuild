package handlers

import (
	"encoding/json"
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ManagerHandler interface {
	ChangeGameStatus(c *fiber.Ctx) error
	SubscribePlayer(c *websocket.Conn)
}

type managerHandler struct {
	service services.ManagerService
}

func NewManagerHandler(service services.ManagerService) ManagerHandler {
	return &managerHandler{
		service: service,
	}
}

// ChangeGameStatus implements GameHandler.
func (g *managerHandler) ChangeGameStatus(c *fiber.Ctx) error {
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

	return c.SendStatus(fiber.StatusOK)
}

// SubscribePlayer implements ManagerHandler.
func (g *managerHandler) SubscribePlayer(c *websocket.Conn) {
	playerID := c.Params("playerID")

	msg, err := g.service.AddPlayer(playerID, c)
	if err != nil {
		g.service.RemovePlayer(playerID)
		c.WriteJSON(fiber.Map{
			"message": "error while adding player",
		})
		return
	}

	if err := c.WriteJSON(msg); err != nil {
		log.Printf("Error writing message to player %s: %v", playerID, err)
		g.service.RemovePlayer(playerID)
		return
	}

	go g.service.HandleBoardcast()

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from player %s: %v", playerID, err)
			g.service.RemovePlayer(playerID)
			break
		}

		if messageType == websocket.TextMessage {
			var text map[string]interface{}
			err := json.Unmarshal(message, &text)
			if err != nil {
				log.Fatalf("Errors : %s", err.Error())
			}
			for key, value := range text {
				log.Printf("%s : '%v'\n", key, value)
			}
			g.service.HandlePlayerMessage(playerID, message)
		}
	}
}
