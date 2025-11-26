// Code generated stub — gRPC interfaces minimal.
package gen

import "context"

type UserServiceServer interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error)
}

// Messages
type CreateUserRequest struct { Name, Email, Password string }
type CreateUserResponse struct { User *User }
type GetUserRequest struct { Id string }
type GetUserResponse struct { User *User }

// Register helper stub
func RegisterUserServiceServer(s interface{}, srv UserServiceServer) {}
