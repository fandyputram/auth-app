package handler

import (
	uc "auth-app/internal/usecase"
	"encoding/json"
	"net/http"
)

// AuthHandler handles HTTP requests for authentication and authorization
type AuthHandler struct {
	authUsecase uc.AuthUsecase
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authUc uc.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUc,
	}
}

// RegisterUserHandler handles HTTP requests for user registration
func (h *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload uc.RegisterPayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.authUsecase.Register(r.Context(), &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// LoginUserHandler handles HTTP requests for user login
func (h *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload uc.LoginPayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.authUsecase.Login(r.Context(), &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Return the logged-in user in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
