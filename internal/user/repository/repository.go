package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/johnqr/user-service/internal/user/domain"
)

var ErrNotFound = errors.New("user not found")
var ErrConflict = errors.New("already exists")

type UserRepository interface {
	Create(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
