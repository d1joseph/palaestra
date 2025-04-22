package api

import (
	"github.com/d1joseph/palaestra/internal/core/ports"
	"github.com/d1joseph/palaestra/internal/domain/models"
	"github.com/gofiber/fiber/v2"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// APIHandler implements the generated ServerInterface
type APIHandler struct {
	userService ports.UserService
	// Add other services as needed
}

// NewAPIHandler creates a new API handler
func NewAPIHandler(userService ports.UserService) *APIHandler {
	return &APIHandler{
		userService: userService,
	}
}

// PostUsers handles user creation
func (h *APIHandler) PostUsers(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.userService.CreateUser(c.Context(), user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("")
}

// GetUsersUserId handles getting a user by ID
func (h *APIHandler) GetUsersUserId(c *fiber.Ctx, userId openapi_types.UUID) error {
	user, err := h.userService.GetUserByID(c.Context(), userId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// PutUsersUserId handles updating a user
func (h *APIHandler) PutUsersUserId(c *fiber.Ctx, userId openapi_types.UUID) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Ensure the ID in the path matches the ID in the body
	user.Id = &userId

	if err := h.userService.UpdateUser(c.Context(), user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("")
}

// GetExercises implementation using api.GetExercisesParams instead of models.GetExercisesParams
func (h *APIHandler) GetExercises(c *fiber.Ctx, params GetExercisesParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

// Additional server methods (implementing ServerInterface)
func (h *APIHandler) PostExercises(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) GetExercisesExerciseId(c *fiber.Ctx, exerciseId openapi_types.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) GetProgress(c *fiber.Ctx, params GetProgressParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) PostProgress(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) GetWorkoutPlans(c *fiber.Ctx, params GetWorkoutPlansParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) PostWorkoutPlans(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) GetWorkouts(c *fiber.Ctx, params GetWorkoutsParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) PostWorkouts(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) GetWorkoutsWorkoutId(c *fiber.Ctx, workoutId openapi_types.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) PutWorkoutsWorkoutId(c *fiber.Ctx, workoutId openapi_types.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) PostWorkoutsWorkoutIdComplete(c *fiber.Ctx, workoutId openapi_types.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *APIHandler) PostWorkoutsWorkoutIdStart(c *fiber.Ctx, workoutId openapi_types.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}
