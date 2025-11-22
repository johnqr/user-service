package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/johnqr/user-service/internal/user/domain"
)

type UserHandler struct {
	service  domain.UserService
	validate *validator.Validate
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, "invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		BadRequest(w, err.Error())
		return
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		BadRequest(w, err.Error())
		return
	}

	resp := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	JSON(w, http.StatusCreated, resp)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.service.GetUser(id)
	if err != nil {
		NotFound(w, err.Error())
		return
	}

	resp := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	JSON(w, http.StatusOK, resp)
}
