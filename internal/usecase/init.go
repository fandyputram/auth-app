package auth

import (
	d "auth-app/internal/domain"
	r "auth-app/internal/interface/repository"
	"context"
)

type AuthUsecase interface {
	Login(context.Context, LoginPayload) (*d.User, error)
	Register(context.Context, RegisterPayload) error
}

type authImpl struct {
	userRepository r.UserRepository
}
