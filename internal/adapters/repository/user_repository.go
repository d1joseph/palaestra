package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/d1joseph/palaestra/internal/core/ports"
	"github.com/d1joseph/palaestra/internal/domain/models"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type memoryUserRepository struct {
	users map[string]models.User
	mu    sync.RWMutex
}

func NewMemoryUserRepository() ports.UserRepository {
	return &memoryUserRepository{
		users: make(map[string]models.User),
	}
}

func (r *memoryUserRepository) CreateUser(ctx context.Context, user models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.Id == nil {
		return errors.New("user ID is required")
	}

	// Check if the user already exists
	if _, exists := r.users[user.Id.String()]; exists {
		return errors.New("user already exists")
	}

	r.users[user.Id.String()] = user
	return nil
}

func (r *memoryUserRepository) GetUserByID(ctx context.Context, id openapi_types.UUID) (models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id.String()]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *memoryUserRepository) UpdateUser(ctx context.Context, user models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.Id == nil {
		return errors.New("user ID is required")
	}

	if _, exists := r.users[user.Id.String()]; !exists {
		return errors.New("user not found")
	}

	r.users[user.Id.String()] = user
	return nil
}
