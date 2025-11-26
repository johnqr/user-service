package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	userhttp "github.com/johnqr/user-service/internal/http/user"
	httpserver "github.com/johnqr/user-service/internal/http"
	"github.com/johnqr/user-service/internal/user/repository"	
	"github.com/johnqr/user-service/internal/user/service"	
)

func main() {
	ctx := context.Background()
	// elegir repo por env (por defecto memoria)
	repo := repository.NewMemoryRepository()
	userSvc := service.NewUserService(repo)
	handler := userhttp.NewUserHandler(userSvc)

	r := httpserver.NewServer(handler)

	addr := ":8080"
	if p := os.Getenv("HTTP_PORT"); p != "" {	addr = ":" + p }

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("HTTP server running on", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	_ = ctx
}
