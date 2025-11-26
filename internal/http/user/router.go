package user

import (
	"github.com/gorilla/mux"
	"github.com/johnqr/user-service/internal/http/middleware"
)

func NewUserRouter(handler *UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Middlewares
	r.Use(middleware.RecoverPanic)
	r.Use(middleware.RequestLogger)

	// Rutas
	r.HandleFunc("/users/register", handler.Register).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")

	return r
}
