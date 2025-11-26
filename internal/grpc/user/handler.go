package user

import (
	"context"
	"github.com/johnqr/user-service/grpc/gen"
	"github.com/johnqr/user-service/internal/user/service"
)

type Handler struct {
	svc service.UserService
}

func NewHandler(s service.UserService) *Handler { return &Handler{svc: s} }

func (h *Handler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	u, err := h.svc.Register(ctx, req.Name, req.Email, req.Password)
	if err != nil { return nil, err }
	return &gen.CreateUserResponse{User: FromDomain(u)}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	u, err := h.svc.GetByID(ctx, req.Id)
	if err != nil { return nil, err }
	return &gen.GetUserResponse{User: FromDomain(u)}, nil
}
