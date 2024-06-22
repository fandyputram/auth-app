package auth

import (
	d "auth-app/internal/domain"
	"context"
)

// RegisterPayload represents the payload structure for user registration
type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Register registers a new user
func (a *authImpl) Register(ctx context.Context, payload *RegisterPayload) error {
	newUser := &d.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}
	return a.userRepository.SaveUser(ctx, newUser)
}
