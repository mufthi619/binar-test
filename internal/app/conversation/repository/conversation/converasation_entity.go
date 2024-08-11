package conversation

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type ConversationEntityGorm struct {
	Id                 uint64        `gorm:"primaryKey"`
	ParticipantUserIds pq.Int64Array `gorm:"type:bigint[]"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (ConversationEntityGorm) TableName() string {
	return "conversations"
}
