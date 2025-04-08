package main

import (
	"context"
	"fmt"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/rest"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/routes"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/manager"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.NewConfig()
	ctx := context.Background()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "connected",
		})
	})

	db := config.ConnectDatabase(cfg, ctx)
	defer db.Close()
	handler := rest.InitHandler(db)

	game := manager.NewGame()
	routes.ManagerRoutes(app, game)
	routes.UserRoutes(app, handler)

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
