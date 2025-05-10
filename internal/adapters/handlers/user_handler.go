package handlers

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetRole(c *fiber.Ctx) error
	UpdateNickname(c *fiber.Ctx) error
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
//
// GetRole handles the request to retrieve a user's role.
//
//	@Summary		Get User Role
//	@Description	Retrieves the role of a user identified by their email.
//	@Tags			Users
//	@Produce		json
//	@Param			email	query		string					true	"User Email"
//	@Success		200		{object}	responses.RoleResponse	"User role retrieved successfully"
//	@Failure		400		{object}	map[string]string		"Bad Request"
//	@Failure		404		{object}	map[string]string		"User Not Found"
//	@Failure		500		{object}	map[string]string		"Internal Server Error"
//	@Router			/user/role [get]
func (u *userHandler) GetRole(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email query parameter is required",
		})
	}
	user, err := u.service.GetUserRole(c.Context(), email)
	switch err {
	case exceptions.ErrUserNotFound:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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

// UpdateNickname implements UserHandler.
// UpdateNickname handles the request to update a user's nickname.
//
//	@Summary		Update User Nickname
//	@Description	Updates the nickname of a user identified by their email.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			email	path		string							true	"User Email"
//	@Param			body	body		requests.ChangeNicknameRequest	true	"Change Nickname Request"
//	@Success		200		{string}	string							"Nickname updated successfully"
//	@Failure		400		{string}	string							"Bad Request"
//	@Failure		500		{object}	map[string]string				"Internal Server Error"
//	@Router			/user/{email}/nickname [put]
func (u *userHandler) UpdateNickname(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var req requests.ChangeNicknameRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := u.service.ChangeUserNickname(c.Context(), email, &req)
	if err != nil {
		switch err {
		case exceptions.ErrEmailNotFound:
			return c.SendStatus(fiber.StatusBadRequest)
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

	}
	return c.SendStatus(fiber.StatusOK)
}
