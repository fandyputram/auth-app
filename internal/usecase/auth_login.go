package usecase

import (
	d "auth-app/internal/domain"
	"context"
)

// LoginPayload represents the payload structure for user login
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login authenticates a user
func (a *authImpl) Login(ctx context.Context, payload *LoginPayload) (*d.User, error) {

	return a.userRepository.GetUserByUsernamePassword(ctx, payload.Username, payload.Password)
}
