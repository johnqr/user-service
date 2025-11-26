package user

import (
	"github.com/johnqr/user-service/grpc/gen"
	"github.com/johnqr/user-service/internal/user/domain"
)

func FromDomain(u *domain.User) *gen.User {
	if u == nil { return nil }
	return &gen.User{Id: u.ID.String(), Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt.String()}
}

func ToDomain(u *gen.User) *domain.User { return &domain.User{} }
