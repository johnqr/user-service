package domain

import (
	"time"

	"github.com/google/uuid"
)

// User representa el modelo de usuario.
type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
