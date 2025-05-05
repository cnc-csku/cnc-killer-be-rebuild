package handlers

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type ActionHandler interface {
	AddAction(c *fiber.Ctx) error
}

type actionHandler struct {
	service services.ActionService
}

func NewActionHandler(service services.ActionService) ActionHandler {
	return &actionHandler{
		service: service,
	}
}

// AddAction handles the HTTP request to add a new action.
// @Summary Add a new action
// @Description Adds a new action with the provided details and condition.
// @Tags Actions
// @Accept json
// @Produce json
// @Param action body models.Action true "Action details"
// @Success 200 {object} map[string]string "Action added successfully"
// @Failure 400 {object} map[string]string "Error message"
// @Router /action [post]
func (a *actionHandler) AddAction(c *fiber.Ctx) error {
	var action models.Action
	if err := c.BodyParser(&action); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := a.service.AddAction(c.Context(), action.Detail, action.Condition); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "action added successfully",
	})
}
