package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewUserRouter(handler *UserHandler) http.Handler {
	r := mux.NewRouter()

	r.Use(RequestLogger)
	r.Use(RecoverPanic)

	r.HandleFunc("/users/register", handler.Register).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")

	return r
}
