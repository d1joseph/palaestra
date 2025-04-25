package api

import (
	"github.com/d1joseph/palaestra/internal/core/ports"
	"github.com/d1joseph/palaestra/internal/domain/models"
	"github.com/gofiber/fiber/v2"
	openapitypes "github.com/oapi-codegen/runtime/types"
)

// Handler implements the generated ServerInterface
type Handler struct {
	userService ports.UserService
	// Add other services as needed
}

// NewAPIHandler creates a new API handler
func NewAPIHandler(userService ports.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

// PostUsers handles user creation
func (h *Handler) PostUsers(c *fiber.Ctx) error {
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
func (h *Handler) GetUsersUserId(c *fiber.Ctx, userId openapitypes.UUID) error {
	user, err := h.userService.GetUserByID(c.Context(), userId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// PutUsersUserId handles updating a user
func (h *Handler) PutUsersUserId(c *fiber.Ctx, userId openapitypes.UUID) error {
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
func (h *Handler) GetExercises(c *fiber.Ctx, params GetExercisesParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

// PostExercises Additional server methods (implementing ServerInterface)
func (h *Handler) PostExercises(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetExercisesExerciseId(c *fiber.Ctx, exerciseId openapitypes.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetProgress(c *fiber.Ctx, params GetProgressParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) PostProgress(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetWorkoutPlans(c *fiber.Ctx, params GetWorkoutPlansParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) PostWorkoutPlans(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetWorkouts(c *fiber.Ctx, params GetWorkoutsParams) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) PostWorkouts(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetWorkoutsWorkoutId(c *fiber.Ctx, workoutId openapitypes.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) PutWorkoutsWorkoutId(c *fiber.Ctx, workoutId openapitypes.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) PostWorkoutsWorkoutIdComplete(c *fiber.Ctx, workoutId openapitypes.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}

func (h *Handler) PostWorkoutsWorkoutIdStart(c *fiber.Ctx, workoutId openapitypes.UUID) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Not implemented")
}
