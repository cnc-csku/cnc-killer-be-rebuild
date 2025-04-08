package routes

import (
	"log"

	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/manager"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ManagerRoutes(app *fiber.App, g *manager.Game) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Post("/game/status", func(c *fiber.Ctx) error {
		var body struct {
			Status string `json:"status"`
		}

		if err := c.BodyParser(&body); err != nil {
			return err
		}

		err := g.ChangeGameStatus(body.Status)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{"status": g.Status})
	})
	api := app.Group("/ws")
	api.Get("/:playerID", websocket.New(func(c *websocket.Conn) {
		playerID := c.Params("playerID")

		g.AddPlayer(playerID, c)

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Error reading message from player %s: %v", playerID, err)
				g.RemovePlayer(playerID)
				break
			}

			if messageType == websocket.TextMessage {
				g.HandlerPlayerMessage(playerID, message)
			}
		}
	}))
	go g.HandlerBoardcast()

}
