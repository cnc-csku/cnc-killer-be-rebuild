package main

import (
	"context"
	"fmt"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
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
	print(db)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
