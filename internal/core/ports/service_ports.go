package ports

import (
	"context"

	"github.com/d1joseph/palaestra/internal/domain/models"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id openapi_types.UUID) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
}
