package repository

import "time"

type NotificationEntityGorm struct {
	Id      uint64    `gorm:"primaryKey;autoIncrement"`
	UserId  uint64    `gorm:"index"`
	Message string    `gorm:"type:varchar(255);not null"`
	SentAt  time.Time `gorm:"index;not null"`
}

func (NotificationEntityGorm) TableName() string {
	return "notifications"
}
