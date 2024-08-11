package conversation

import (
	"binar/internal/app/conversation/domain"
	"github.com/lib/pq"
)

func ToConversationEntityGorm(data domain.Conversation) ConversationEntityGorm {
	var pqArr pq.Int64Array
	for _, val := range data.Participants {
		pqArr = append(pqArr, int64(val))
	}
	return ConversationEntityGorm{
		Id:                 data.Id,
		ParticipantUserIds: pqArr,
		CreatedAt:          data.CreatedAt,
	}
}

func ToConversationDomain(entity ConversationEntityGorm) domain.Conversation {
	var nativeArr []uint64
	for _, val := range entity.ParticipantUserIds {
		nativeArr = append(nativeArr, uint64(val))
	}
	return domain.Conversation{
		Id:           entity.Id,
		Participants: nativeArr,
		CreatedAt:    entity.CreatedAt,
	}
}

func ToConversationsDomain(entities []ConversationEntityGorm) (finalResponse []domain.Conversation) {
	for _, val := range entities {
		finalResponse = append(finalResponse, ToConversationDomain(val))
	}
	return
}
