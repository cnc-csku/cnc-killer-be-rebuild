package handlers

import (
	"errors"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/requests"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/gofiber/fiber/v2"
)

type ActionHandler interface {
	AddAction(c *fiber.Ctx) error
	FindActionByID(c *fiber.Ctx) error
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
// @Param req body requests.AddActionRequest true "Add Action Request"
// @Success 201 {object} map[string]interface{} "Success response with action details"
// @Failure 400 {object} map[string]interface{} "Invalid action data provided"
// @Failure 500 {object} map[string]interface{} "Failed to add an action"
// @Router /action [post]
func (a *actionHandler) AddAction(c *fiber.Ctx) error {
	var req requests.AddActionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if req.Detail == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Detail field is required",
		})
	}

	if err := a.service.AddAction(c.Context(), &req); err != nil {
		switch {
		case errors.Is(err, exceptions.ErrInvalidAction):
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":   "error",
				"messsage": "Invalid action data provided",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to add an action",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   req,
	})
}

// FindActionByID handles the HTTP request to retrieve an action by its ID.
// @Summary      Retrieve an action by ID
// @Description  Fetches an action from the database using the provided ID.
// @Tags         Actions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Action ID"
// @Success      200  {object}  map[string]interface{}  "Action retrieved successfully"
// @Failure      400  {object}  map[string]interface{}  "Invalid action ID provided"
// @Failure      404  {object}  map[string]interface{}  "Action not found"
// @Failure      500  {object}  map[string]interface{}  "Internal server error"
// @Router       /action/{id} [get]
func (a *actionHandler) FindActionByID(c *fiber.Ctx) error {
	id := c.Params("id")

	action, err := a.service.FindActionByID(c.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrEmptyActionID):
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid action ID provided",
			})
		case errors.Is(err, exceptions.ErrActionNotFound):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "error",
				"error":  "Action not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   action,
	})
}
