package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/rest"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ManagerRoutes(app *fiber.App, handler *rest.Handler) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Post("/game/status", handler.ManagerHandler.ChangeGameStatus)
	api := app.Group("/ws")
	api.Get("/:playerID", websocket.New(handler.ManagerHandler.SubscribePlater))

}
