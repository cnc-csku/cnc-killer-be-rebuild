package rest

import (
	"context"
	"io"
	"net/http"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/gofiber/fiber/v2"
)

type GoogleAuthHandler interface {
	GoogleLogin(c *fiber.Ctx) error
	GoogleCallback(c *fiber.Ctx) error
}

type googleAuthHandler struct {
	GoogleCfg *config.GoogleAuthConfig
}

func NewGoogleAuthHandler(cfg *config.GoogleAuthConfig) GoogleAuthHandler {
	return googleAuthHandler{
		GoogleCfg: cfg,
	}
}

// GoogleLogin implements GoogleAuthHandler.
func (g googleAuthHandler) GoogleLogin(c *fiber.Ctx) error {
	url := g.GoogleCfg.AuthConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)

	return c.JSON(url)
}

// GoogleCallback implements GoogleAuthHandler.
func (g googleAuthHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("State don't Match")
	}

	ctx := context.Background()
	code := c.Query("code")

	token, err := g.GoogleCfg.AuthConfig.Exchange(ctx, code)

	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	defer resp.Body.Close()

	user, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(user))

}
