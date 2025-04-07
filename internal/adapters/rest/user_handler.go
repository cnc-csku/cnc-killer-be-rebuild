package rest

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetRole(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
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
	var req requests.UserGetRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := u.service.GetUserRole(c.Context(), req.UserID)
	switch err {
	case exceptions.ErrUserNotFound:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
		})
	case exceptions.ErrInvalidUUID:
		return c.SendStatus(fiber.StatusBadRequest)
	case nil:
		return c.Status(fiber.StatusOK).JSON(user)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}

// Login implements UserHandler.
func (u *userHandler) Login(c *fiber.Ctx) error {
	var req requests.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user, err := u.service.Login(c.Context(), req)
	switch err {
	case exceptions.ErrUnauthorized:
		return c.SendStatus(fiber.StatusUnauthorized)
	case nil:
		return c.Status(fiber.StatusOK).JSON(user)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}
