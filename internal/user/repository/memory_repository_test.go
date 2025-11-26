package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/johnqr/user-service/internal/user/domain"
	"github.com/johnqr/user-service/internal/user/repository"
)

func TestMemoryCreateGet(t *testing.T) {
	m := repository.NewMemoryRepository()
	ctx := context.Background()

	// Crear un usuario con ID asignado
	u := &domain.User{
		ID:    uuid.New(),
		Name:  "John",
		Email: "john@example.com",
	}

	// Guardar el usuario
	err := m.Create(ctx, u)
	if err != nil {
		t.Fatalf("unexpected error on Create: %v", err)
	}

	// Obtener por ID (uuid.UUID, no string)
	got, err := m.GetByID(ctx, u.ID)
	if err != nil {
		t.Fatalf("unexpected error on GetByID: %v", err)
	}

	if got.Email != u.Email {
		t.Errorf("expected email %s, got %s", u.Email, got.Email)
	}
}
