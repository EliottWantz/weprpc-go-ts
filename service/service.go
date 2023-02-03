package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type ServiceRPC struct {
	users map[string]*User
	mu    sync.Mutex
}

var _ UserService = (*ServiceRPC)(nil)

func NewUserService() *ServiceRPC {
	return &ServiceRPC{
		users: make(map[string]*User),
	}
}

func (s *ServiceRPC) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *ServiceRPC) CreateUser(ctx context.Context, username string, password string) (*User, error) {
	if username == "" {
		return nil, WrapError(ErrInvalidArgument, errors.New("username is empty"), "username is empty")
	}
	if password == "" {
		return nil, WrapError(ErrInvalidArgument, errors.New("password is empty"), "password is empty")
	}

	u := &User{
		Username:  username,
		Password:  password,
		ID:        fmt.Sprint(len(s.users) + 1),
		CreatedAt: time.Now(),
	}

	s.mu.Lock()
	s.users[u.ID] = u
	s.mu.Unlock()

	return u, nil
}

func (s *ServiceRPC) GetUserByID(ctx context.Context, userID string) (*User, error) {
	if userID == "" {
		return nil, WrapError(ErrInvalidArgument, errors.New("userID is empty"), "userID is empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func (s *ServiceRPC) ListUsers(ctx context.Context) ([]*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]*User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}

	return users, nil
}
