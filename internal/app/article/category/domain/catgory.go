package domain

import "time"

type Category struct {
	Id        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type CategoryRepository interface {
	GetAll() ([]Category, error)
}

type CategoryService interface {
	GetAll() ([]Category, string, error)
}
