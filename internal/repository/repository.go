package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UsersRepository Users
}

type (
	Users interface {
		Create(*User) error
		Update(uuid.UUID, *User) (*User, error)
		GetByID(uuid.UUID) (*User, error)
		GetAll(*Pagination) []User
		DeleteByID(uuid.UUID) (*User, error)
	}
)

type (
	Pagination struct {
		Page int
		Size int
	}
)

func NewRepositories(db *sqlx.DB) *Repositories {
	usersRepository := NewUsersRepo(db)

	return &Repositories{UsersRepository: usersRepository}
}
