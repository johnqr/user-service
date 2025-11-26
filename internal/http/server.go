package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnqr/user-service/internal/http/middleware"
	userhttp "github.com/johnqr/user-service/internal/http/user"
)

func NewServer(handler *userhttp.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.RecoverPanic)
	r.Use(middleware.RequestLogger)

	r.PathPrefix("/users").Handler(userhttp.NewUserRouter(handler))

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}).Methods("GET")

	return r
}
