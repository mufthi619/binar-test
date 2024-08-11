package domain

import "time"

type Message struct {
	Id             uint64
	ConversationId uint64
	SenderId       uint64
	Content        string
	SentAt         time.Time
}

type Conversation struct {
	Id           uint64
	Participants []uint64
	CreatedAt    time.Time
}

type MessageRepository interface {
	Create(data Message) (*Message, error)
	GetAllByConversationId(conversationId uint64) ([]Message, error)
}

type ConversationRepository interface {
	Create(data Conversation) (*Conversation, error)
	GetById(id uint64) (*Conversation, error)
}

type MessageService interface {
	Create(data Message) (*Message, string, error)
	GetAllByConversationId(conversationId uint64) ([]Message, string, error)
}

type ConversationService interface {
	Create(participants []uint64) (*Conversation, string, error)
	GetById(id uint64) (*Conversation, string, error)
}
