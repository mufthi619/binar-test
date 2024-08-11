package message

import "time"

type MessageEntityGorm struct {
	Id             uint64    `gorm:"primaryKey;autoIncrement"`
	ConversationId uint64    `gorm:"index;not null"`
	SenderId       uint64    `gorm:"index;not null"`
	Content        string    `gorm:"type:text;not null"`
	SentAt         time.Time `gorm:"index;not null"`
}

func (MessageEntityGorm) TableName() string {
	return "messages"
}
