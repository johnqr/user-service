package service_test

import (
	"context"
	"testing"

	"github.com/johnqr/user-service/internal/user/repository"
	"github.com/johnqr/user-service/internal/user/service"
)

func TestRegisterAndDuplicate(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewUserService(repo)
	ctx := context.Background()
	u, err := svc.Register(ctx, "John", "a@b.com", "password123")
	if err != nil {
		t.Fatal(err)
	}
	if u.Email != "a@b.com" {
		t.Fatalf("email mismatch")
	}
	// duplicate
	if _, err := svc.Register(ctx, "X", "a@b.com", "password123"); err == nil {
		t.Fatal("expected error on duplicate")
	}
}
