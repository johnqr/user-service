package http

import (
	"github.com/gorilla/mux"
)

func NewUserRouter(handler *UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.Use(RequestLogger)
	r.Use(RecoverPanic)

	r.HandleFunc("/users/register", handler.Register).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")

	return r
}
