package repository

import (
	"errors"
	"sync"

	"github.com/johnqr/user-service/internal/user/domain"
)

// MemoryRepository implementa UserRepository pero en memoria.
// Es ideal para pruebas, desarrollo sin base de datos y unit tests.
type MemoryRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

// NewMemoryRepository crea un repositorio en memoria inicializado.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[string]*domain.User),
	}
}

// Create agrega un nuevo usuario al repositorio.
func (r *MemoryRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verificar si ya existe un email igual
	for _, u := range r.users {
		if u.Email == user.Email {
			return errors.New("email ya está registrado")
		}
	}

	r.users[user.ID] = user
	return nil
}

// GetByID obtiene un usuario por su ID.
func (r *MemoryRepository) GetByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("usuario no encontrado")
	}

	return user, nil
}

// GetByEmail obtiene un usuario por su correo.
func (r *MemoryRepository) GetByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}

// Update actualiza un usuario existente.
func (r *MemoryRepository) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.users[user.ID]
	if !exists {
		return errors.New("usuario no encontrado")
	}

	r.users[user.ID] = user
	return nil
}

// Delete elimina un usuario por ID.
func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.users[id]
	if !exists {
		return errors.New("usuario no encontrado")
	}

	delete(r.users, id)
	return nil
}
