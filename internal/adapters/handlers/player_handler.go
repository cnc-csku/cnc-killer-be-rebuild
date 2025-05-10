package handlers

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type PlayerHandler interface {
	AddPlayer(c *fiber.Ctx) error
}

type playerHandler struct {
	service services.PlayerService
}

func NewPlayerHandler(service services.PlayerService) PlayerHandler {
	return &playerHandler{
		service: service,
	}
}

// AddPlayer godoc
// @Summary Add a new player
// @Description Add a new player using the provided JWT token
// @Tags Players
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /player [post]
func (p *playerHandler) AddPlayer(c *fiber.Ctx) error {
	token := c.Cookies("token")

	if err := p.service.AddPlayer(c.Context(), token); err != nil {
		switch err {
		case exceptions.ErrNicknameNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Nickname not found",
			})
		case exceptions.ErrTokenIsEmpty:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Token is empty",
			})
		case exceptions.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "An unexpected error occurred",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Player added successfully",
	})
}
