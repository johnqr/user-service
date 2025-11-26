package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/johnqr/user-service/internal/user/domain"
)

// MemoryRepository is an in-memory implementation of UserRepository.
// Useful for development and unit tests.
type MemoryRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*domain.User
}

// NewMemoryRepository creates a new in-memory repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[uuid.UUID]*domain.User),
	}
}

// Create stores a new user in memory.
func (m *MemoryRepository) Create(ctx context.Context, u *domain.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	// Prevent duplicate emails.
	for _, existing := range m.users {
		if existing.Email == u.Email {
			return errors.New("email already exists")
		}
	}

	m.users[u.ID] = u
	return nil
}

// GetByID retrieves a user by ID.
func (m *MemoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	u, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}

// GetByEmail retrieves a user by email.
func (m *MemoryRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

// Delete removes a user by ID.
func (m *MemoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(m.users, id)
	return nil
}

func (m *MemoryRepository) Update(ctx context.Context, u *domain.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.users[u.ID]; !exists {
		return errors.New("user not found")
	}

	m.users[u.ID] = u
	return nil
}
