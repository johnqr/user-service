package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnqr/user-service/internal/user/service"
)

// UserHandler maneja endpoints de usuario.
type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler { return &UserHandler{svc: s} }

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	u, err := h.svc.Register(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		switch err {
		case service.ErrInvalidEmail:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrWeakPassword:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrUserExists:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "internal", http.StatusInternalServerError)
		}
		return
	}
	resp := RegisterResponse{ID: u.ID.String(), Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt}
	w.Header().Set("Location", "/users/"+u.ID.String())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	u, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(RegisterResponse{ID: u.ID.String(), Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt})
}
