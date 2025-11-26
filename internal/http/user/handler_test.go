package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	userhttp "github.com/johnqr/user-service/internal/http/user"
	"github.com/johnqr/user-service/internal/user/repository"
	"github.com/johnqr/user-service/internal/user/service"
)

func TestRegisterHandler(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewUserService(repo)
	h := userhttp.NewUserHandler(svc)
	reqBody, _ := json.Marshal(map[string]string{"name":"J","email":"x@y.com","password":"password123"})
	r := httptest.NewRequest("POST", "/users/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	h.Register(w, r)
	if w.Code != http.StatusCreated { t.Fatalf("expected 201 got %d", w.Code) }
}
