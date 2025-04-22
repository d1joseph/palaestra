package ports

import (
	"context"

	"github.com/d1joseph/palaestra/internal/domain/models"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// UserRepository defines the interface for user data persistence
type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id openapi_types.UUID) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
}
