package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	userhttp "github.com/johnqr/user-service/internal/http/user"
	"github.com/johnqr/user-service/internal/user/repository"
	"github.com/johnqr/user-service/internal/user/service"
)

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Leer puerto desde variables de entorno (útil para despliegues)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Dependencias
	repo := repository.NewMemoryRepository()
	userService := service.NewUserService(repo)
	handler := userhttp.NewUserHandler(userService)

	// Router
	router := userhttp.NewUserRouter(handler)
	router.Use(loggingMiddleware)

	// Configurar servidor HTTP profesional
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Canal para señales del sistema (Ctrl + C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Iniciar servidor
	go func() {
		log.Printf("🚀 Servidor HTTP corriendo en http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Error al iniciar servidor: %v", err)
		}
	}()

	// Esperar señal de cierre
	<-stop
	log.Println("⏳ Cerrando servidor...")

	// Contexto para apagado elegante
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Error durante shutdown: %v", err)
	}

	log.Println("✅ Servidor detenido correctamente")
}
