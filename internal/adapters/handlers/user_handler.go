package handlers

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetRole(c *fiber.Ctx) error
}
type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

// GetRole implements UserHandler.
func (u *userHandler) GetRole(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user, err := u.service.GetUserRole(c.Context(), email)
	switch err {
	case exceptions.ErrUserNotFound:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
		})
	case nil:
		return c.Status(fiber.StatusOK).JSON(user)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}
