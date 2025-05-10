package handlers

import (
	"time"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type GoogleAuthHandler interface {
	GoogleLogin(c *fiber.Ctx) error
	GoogleCallback(c *fiber.Ctx) error
	GetRefreshToken(c *fiber.Ctx) error
}

type googleAuthHandler struct {
	service services.AuthService
}

func NewGoogleAuthHandler(service services.AuthService) GoogleAuthHandler {
	return &googleAuthHandler{
		service: service,
	}
}

// GoogleCallback implements GoogleAuthHandler.
func (g *googleAuthHandler) GoogleCallback(c *fiber.Ctx) error {
	user, err := g.service.GetUserInfo(c)
	// tokenStr := c.Cookies("token")
	// token, err := utils.JWTDecode(tokenStr)
	// log.Print(token)
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

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   user.Token,
		Expires: time.Now().Add(5 * time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(user)

}

// GoogleLogin implements GoogleAuthHandler.

//	@Tags			Auth
//	@Summary		login
//	@Description	login with google
//	@Success		200
//	@Router			/auth/google [get]
func (g *googleAuthHandler) GoogleLogin(c *fiber.Ctx) error {
	authURL, err := g.service.GetAuthURL()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There's some error while get sign in url",
		})
	}

	c.Status(fiber.StatusSeeOther)

	return c.Redirect(authURL)
}

// GetRefreshToken implements GoogleAuthHandler.
func (g *googleAuthHandler) GetRefreshToken(c *fiber.Ctx) error {
	var req requests.TokenRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := g.service.GetRefreshToken(c.Context(), &req)
	if err != nil {
		switch err {
		case exceptions.ErrEmailNotFound:
		case exceptions.ErrUserNotFound:
			return c.SendStatus(fiber.StatusBadRequest)
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
	})
}
