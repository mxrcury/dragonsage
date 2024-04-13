package service

import (
	"context"

	"github.com/mxrcury/dragonsage/internal/repository"
)

type Services struct {
	UsersService *UsersService
}

type (
	Users interface {
		SignUp(context.Context) error
	}
)

func NewServices(repos *repository.Repositories) *Services {
	usersService := NewUsersService(repos.UsersRepository)

	return &Services{UsersService: usersService}
}
