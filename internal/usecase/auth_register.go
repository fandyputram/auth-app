package usecase

import (
	d "auth-app/internal/domain"
	"auth-app/internal/utils"
	"context"
	"errors"
)

// RegisterPayload represents the payload structure for user registration
type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Register registers a new user
func (a *authImpl) Register(ctx context.Context, payload *RegisterPayload) error {
	err := validateRegisterPayload(payload)
	if err != nil {
		return err
	}

	newUser := &d.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}
	return a.userRepository.SaveUser(ctx, newUser)
}

func validateRegisterPayload(payload *RegisterPayload) error {
	if payload.Email == "" || payload.Password == "" || payload.Username == "" {
		return errors.New("empty payload")
	}

	if !utils.ValidateEmailFormat(payload.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
