package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxrcury/dragonsage/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

type UsersSignUpInput struct {
	Username string
	Password string
	Email    string
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) SignUp(input *UsersSignUpInput) error {
	user := &repository.User{
		ID:        uuid.New(),
		Username:  input.Username,
		Password:  input.Password,
		Email:     input.Email,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
	}

	err := s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
