package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func ActionRoutes(app *fiber.App, handler *handlers.Handler) {
	api := app.Group("/action")
	api.Post("/", handler.ActionHandler.AddAction)
	api.Get("/:id", handler.ActionHandler.FindActionByID)
}
