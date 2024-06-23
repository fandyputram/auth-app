package usecase

import (
	d "auth-app/internal/domain"
	r "auth-app/internal/interface/repository"
	"context"
)

type AuthUsecase interface {
	Login(context.Context, *LoginPayload) (*d.User, error)
	Register(context.Context, *RegisterPayload) error
}

type authImpl struct {
	userRepository r.UserRepository
}

// NewAuthUseCase creates a new instance of authUseCase
func NewAuthUseCase(userRepo r.UserRepository) AuthUsecase {
	return &authImpl{
		userRepository: userRepo,
	}
}
