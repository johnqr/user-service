package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/johnqr/user-service/internal/user/domain"
	"github.com/johnqr/user-service/internal/user/repository"
)

var (
	ErrInvalidEmail = errors.New("email inválido")
	ErrWeakPassword = errors.New("password débil")
	ErrUserExists = repository.ErrConflict
	ErrUserNotFound = repository.ErrNotFound
)

// UserService define la lógica de usuarios.
type UserService interface {
	Register(ctx context.Context, name, email, password string) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Register(ctx context.Context, name, email, password string) (*domain.User, error) {
	// validar email sencillo
	re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	if !re.MatchString(email) { return nil, ErrInvalidEmail }
	if len(password) < 8 { return nil, ErrWeakPassword }
	// comprobar existencia
	if _, err := s.repo.GetByEmail(ctx, email); err == nil {
		return nil, ErrUserExists
	}
	// hash
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil { return nil, err }
	u := &domain.User{
		ID: uuid.New(),
		Name: name,
		Email: email,
		PasswordHash: string(h),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	if err := s.repo.Create(ctx, u); err != nil { return nil, err }
	return u, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil { return nil, ErrUserNotFound }
	u, err := s.repo.GetByID(ctx, uid)
	if err != nil { return nil, ErrUserNotFound }
	return u, nil
}
