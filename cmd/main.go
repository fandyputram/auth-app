package main

import (
	"auth-app/internal/interface/handler"
	"auth-app/internal/interface/repository"
	"auth-app/internal/usecase"
	"context"
	"log"
	"net/http"
)

func main() {
	// Initialize repositories
	userRepository := repository.NewUserRepository()

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepository)

	// Initialize HTTP handlers
	authHandler := handler.NewAuthHandler(authUseCase)

	// Register HTTP handlers with context support
	http.HandleFunc("/register", withContext(authHandler.RegisterUserHandler))
	http.HandleFunc("/login", withContext(authHandler.LoginUserHandler))

	// Start HTTP server
	log.Println("running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// withContext is a middleware function to inject context into HTTP handlers
func withContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// Add any context values here if needed

		// Call the next handler with the request context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
