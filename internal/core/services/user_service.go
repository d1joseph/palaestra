package services

import (
	"context"
	"fmt"
	"time"

	"github.com/d1joseph/palaestra/internal/core/ports"
	"github.com/d1joseph/palaestra/internal/domain/models"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user models.User) error {

	// Generate a user ID if not provided
	if user.Id == nil {
		// Create a UUID directly
		id := openapi_types.UUID(uuid.New())
		user.Id = &id
	}

	fmt.Printf("Generating new user ID: %v\n", *user.Id)

	// Set timestamps
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now

	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id openapi_types.UUID) (models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, user models.User) error {
	// Update timestamp
	now := time.Now()
	user.UpdatedAt = &now

	return s.userRepo.UpdateUser(ctx, user)
}
