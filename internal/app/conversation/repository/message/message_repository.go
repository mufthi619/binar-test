package message

import (
	"binar/internal/app/conversation/domain"
	"binar/internal/infra/database"
	"go.uber.org/zap"
)

type messageRepository struct {
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewMessageRepository(databaseClient *database.Databases, logger *zap.Logger) domain.MessageRepository {
	return &messageRepository{
		databaseClient: databaseClient,
		logger:         logger,
	}
}

func (r *messageRepository) Create(data domain.Message) (*domain.Message, error) {
	entity := ToMessageEntityGorm(data)

	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}
	if err := writer.Create(&entity).Error; err != nil {
		return nil, err
	}
	finalResponse := ToMessageDomain(entity)

	return &finalResponse, nil
}

func (r *messageRepository) GetAllByConversationId(conversationId uint64) ([]domain.Message, error) {
	reader := r.databaseClient.ReadDB

	var resp []MessageEntityGorm
	err := reader.Model(&MessageEntityGorm{}).
		Where("conversation_id = ?", conversationId).
		Order("sent_at ASC").
		Scan(&resp).Error
	if err != nil {
		return nil, err
	}
	finalResponse := ToMessagesDomain(resp)

	return finalResponse, nil
}
