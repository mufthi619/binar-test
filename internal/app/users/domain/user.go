package domain

import (
	"time"
)

type User struct {
	ID        uint64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserRepository interface {
	Create(data User) (*User, error)
	GetById(id uint64) (*User, error)
	GetByIdIn(id []uint64) ([]User, error)
}

type UserService interface {
	Create(data User) (*User, string, error)
	GetById(id uint64) (*User, string, error)
}
