package main

import (
	"log"
	"net/http"

	userhttp "github.com/johnqr/user-service/internal/http"
	"github.com/johnqr/user-service/internal/user/repository"
	"github.com/johnqr/user-service/internal/user/service"
)

func main() {
	repo := repository.NewMemoryRepository()
	userService := service.NewUserService(repo)
	handler := userhttp.NewUserHandler(userService)

	router := userhttp.NewUserRouter(handler)

	log.Println("HTTP server running on :8080")
	http.ListenAndServe(":8080", router)
}
