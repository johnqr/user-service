package domain

import "time"

// User representa la entidad central de nuestro dominio.
// Es el modelo base que usará el servicio, repositorio y los handlers.
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository define las operaciones que cualquier repositorio debe implementar.
// Puede haber implementaciones como PostgresUserRepository, MemoryUserRepository, MongoUserRepository, etc.
type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

// UserService define la lógica de negocio del usuario.
// Aquí se definen las reglas, no detalles de infraestructura.
type UserService interface {
	Register(name, email, password string) (*User, error)
	GetUser(id string) (*User, error)
}
