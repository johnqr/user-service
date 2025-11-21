package service

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/johnqr/user-service/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
)

// userService implementa la interfaz UserService del dominio.
// Aquí vive la lógica de negocio REAL.
type userService struct {
	repo domain.UserRepository
}

// NewUserService devuelve una nueva instancia del servicio de usuarios.
// Esto permite inyectar cualquier repositorio que implemente la interfaz.
func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

// Register crea un nuevo usuario aplicando reglas de negocio:
// - Validar nombre
// - Validar email
// - Validar contraseña
// - Verificar si ya existe el usuario
// - Hashear contraseña
// - Generar UUID
// - Crear timestamps
func (s *userService) Register(name, email, password string) (*domain.User, error) {

	// Validación básica
	if len(name) < 3 {
		return nil, errors.New("el nombre es muy corto")
	}

	email = strings.TrimSpace(strings.ToLower(email))
	if !strings.Contains(email, "@") {
		return nil, errors.New("email inválido")
	}

	if len(password) < 6 {
		return nil, errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	// Verificar si el email ya existe
	exists, _ := s.repo.GetByEmail(email)
	if exists != nil {
		return nil, errors.New("el email ya está registrado")
	}

	// Hashear contraseña
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error al encriptar la contraseña")
	}

	// Crear usuario
	user := &domain.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  string(hashedPass),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Guardar en repositorio
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser obtiene un usuario por ID
func (s *userService) GetUser(id string) (*domain.User, error) {

	if id == "" {
		return nil, errors.New("ID vacío")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	return user, nil
}
