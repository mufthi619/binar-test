package payload

import "time"

type (
	MessageDetailResponse struct {
		Id             uint64    `json:"id"`
		ConversationId uint64    `json:"conversation_id"`
		SenderId       uint64    `json:"sender_id"`
		Content        string    `json:"content"`
		SentAt         time.Time `json:"sent_at"`
	}

	CreateConversationRequest struct {
		Participants []uint64 `json:"participants" validate:"required,min=2,dive,required"`
	}

	FindConversation struct {
		Id uint64 `json:"id" query:"id" param:"id"`
	}
	ConversationDetailResponse struct {
		Id           uint64    `json:"id"`
		Participants []uint64  `json:"participants"`
		CreatedAt    time.Time `json:"created_at"`
	}
)
