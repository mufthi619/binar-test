package repository

import "time"

type CategoryEntityGorm struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"unique;not null;size:255"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

func (CategoryEntityGorm) TableName() string {
	return "categories"
}
