package repository

import "time"

type FileEntityGorm struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement"`
	UserId     uint64    `gorm:"index"`
	FileUrl    string    `gorm:"type:varchar(255);not null"`
	UploadedAt time.Time `gorm:"index;not null"`
}

func (FileEntityGorm) TableName() string {
	return "files"
}
