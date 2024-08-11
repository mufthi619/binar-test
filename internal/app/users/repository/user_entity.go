package repository

import (
	"time"
)

type UserEntityGorm struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string     `json:"username" gorm:"unique;not null;size:255"`
	Email     string     `json:"email" gorm:"unique;not null;size:255"`
	Password  string     `json:"password" gorm:"not null;size:255"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

func (UserEntityGorm) TableName() string {
	return "users"
}
