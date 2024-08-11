package repository

import "time"

type JobEntityGorm struct {
	Id          uint64     `gorm:"primaryKey;autoIncrement"`
	Status      string     `gorm:"type:varchar(20);not null;index"`
	QueuedAt    time.Time  `gorm:"not null;index"`
	CompletedAt *time.Time `gorm:"index"`
	Message     string     `gorm:"type:varchar(255);not null"`
}

func (JobEntityGorm) TableName() string {
	return "jobs"
}
