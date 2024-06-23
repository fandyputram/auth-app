package usecase

import (
	d "auth-app/internal/domain"
	"context"
	"errors"
)

// LoginPayload represents the payload structure for user login
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login authenticates a user
func (a *authImpl) Login(ctx context.Context, payload *LoginPayload) (*d.User, error) {
	err := validateLoginPayload(payload)
	if err != nil {
		return nil, err
	}

	return a.userRepository.GetUserByUsernamePassword(ctx, payload.Username, payload.Password)
}

func validateLoginPayload(payload *LoginPayload) error {
	if payload.Password == "" || payload.Username == "" {
		return errors.New("empty payload")
	}

	return nil
}
