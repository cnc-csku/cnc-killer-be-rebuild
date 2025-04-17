package main

import (
	"context"
	"fmt"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/rest"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.NewConfig()
	ctx := context.Background()
	app := fiber.New()

	db := config.ConnectDatabase(cfg, ctx)
	googleCfg := config.NewGoogleConfig(cfg)

	defer db.Close()

	handler := rest.InitHandler(db, googleCfg)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "connected",
		})
	})

	routes.ManagerRoutes(app, handler)
	routes.UserRoutes(app, handler)
	routes.AuthRoute(app, handler)

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
