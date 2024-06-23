package main

import (
	"auth-app/internal/interface/handler"
	"auth-app/internal/interface/repository"
	"auth-app/internal/usecase"
	"context"
	"log"
	"net/http"
	"time"
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

	// Define server configuration
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil, // using the default http.DefaultServeMux
	}

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port 8080: %v\n", err)
		}
	}()

	// Wait briefly to ensure server starts
	time.Sleep(100 * time.Millisecond)
	log.Println("Server is running on port 8080")
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
