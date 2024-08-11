package message

import "binar/internal/app/conversation/domain"

func ToMessageEntityGorm(data domain.Message) MessageEntityGorm {
	return MessageEntityGorm{
		Id:             data.Id,
		ConversationId: data.ConversationId,
		SenderId:       data.SenderId,
		Content:        data.Content,
		SentAt:         data.SentAt,
	}
}

func ToMessageDomain(entity MessageEntityGorm) domain.Message {
	return domain.Message{
		Id:             entity.Id,
		ConversationId: entity.ConversationId,
		SenderId:       entity.SenderId,
		Content:        entity.Content,
		SentAt:         entity.SentAt,
	}
}

func ToMessagesDomain(entities []MessageEntityGorm) (finalResponse []domain.Message) {
	for _, val := range entities {
		resp := domain.Message{
			Id:             val.Id,
			ConversationId: val.ConversationId,
			SenderId:       val.SenderId,
			Content:        val.Content,
			SentAt:         val.SentAt,
		}
		finalResponse = append(finalResponse, resp)
	}
	return
}
