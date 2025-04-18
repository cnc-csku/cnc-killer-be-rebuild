package rest

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type GoogleAuthHandler interface {
	GoogleLogin(c *fiber.Ctx) error
	GoogleCallback(c *fiber.Ctx) error
}

type googleAuthHandler struct {
	service services.AuthService
}

func NewGoogleAuthHandler(service services.AuthService) GoogleAuthHandler {
	return googleAuthHandler{
		service: service,
	}
}

// GoogleCallback implements GoogleAuthHandler.
func (g googleAuthHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")
	user, err := g.service.GetUserInfo(requests.GoogleAuthInfoRequest{
		Code:  code,
		State: state,
	}, c.Context())

	if err != nil {
		switch err {
		case exceptions.ErrCodeNotFound:
		case exceptions.ErrNoState:
			return c.SendStatus(fiber.StatusInternalServerError)
		case exceptions.ErrInvalidState:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "State Doesn't Match",
			})
		case exceptions.ErrExchangeFailed:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Exchange Failed",
			})
		case exceptions.ErrFetchGoogleUser:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while fetching user data",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}

	}

	return c.Status(fiber.StatusOK).JSON(user)

}

// GoogleLogin implements GoogleAuthHandler.
func (g googleAuthHandler) GoogleLogin(c *fiber.Ctx) error {
	authURL, err := g.service.GetAuthURL()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There's some error while get sign in url",
		})
	}

	c.Status(fiber.StatusSeeOther)

	return c.Redirect(authURL)
}
