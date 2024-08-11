package payload

type (
	CreateMessageRequest struct {
		ConversationId uint64 `json:"conversation_id" param:"conversation_id" validate:"required"`
		SenderId       uint64 `json:"sender_id" validate:"required"`
		Content        string `json:"content" validate:"required,min=1,max=1000"`
	}
	FindMessage struct {
		Id uint64 `json:"id" query:"id" param:"id"`
	}
	FindMessagesByConversation struct {
		ConversationId uint64 `json:"conversation_id" query:"conversation_id" param:"conversation_id"`
	}
)
