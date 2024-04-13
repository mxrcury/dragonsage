package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`

	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) Users {
	return &UsersRepo{db: db}
}

func (u *UsersRepo) Create(user *User) error {
  err := u.db.Exec(""bbb
}

func (u *UsersRepo) Update(id uuid.UUID, user *User) (*User, error) {
	return nil, nil
}

func (u *UsersRepo) GetByID(id uuid.UUID) (*User, error) {
	return nil, nil
}

func (u *UsersRepo) GetAll(pagination *Pagination) []User {
	return nil
}
func (u *UsersRepo) DeleteByID(id uuid.UUID) (*User, error) {
	return nil, nil
}
