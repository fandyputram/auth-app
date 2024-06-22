package repository

import (
	d "auth-app/internal/domain"
	"context"
	"errors"
	"fmt"
	"sync"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *d.User) error
	GetUserByUsernamePassword(ctx context.Context, username, password string) (*d.User, error)
}

type userImpl struct {
	users   map[string]*d.User
	counter int
	mu      sync.RWMutex
}

func NewUserRepository() UserRepository {
	return &userImpl{
		users:   make(map[string]*d.User),
		counter: 1,
	}
}

// Save saves a user to memory
func (u *userImpl) SaveUser(ctx context.Context, user *d.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.users[user.Username]; ok {
		return errors.New("username already exists")
	}

	user.ID = u.generateID()

	u.users[user.Username] = user
	return nil
}

// FindByUsernameAndPassword finds a user by username and password
func (u *userImpl) GetUserByUsernamePassword(ctx context.Context, username, password string) (*d.User, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	user, ok := u.users[username]
	if !ok || user.Password != password {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

// generateID generates a simple unique ID (for demonstration purposes)
func (u *userImpl) generateID() string {
	generatedID := fmt.Sprintf("ID%d", u.counter)
	u.counter++
	return generatedID
}
