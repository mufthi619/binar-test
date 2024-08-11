package conversation

import (
	"binar/internal/app/conversation/domain"
	"binar/internal/infra/database"
	"go.uber.org/zap"
)

type conversationRepository struct {
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewConversationRepository(databaseClient *database.Databases, logger *zap.Logger) domain.ConversationRepository {
	return &conversationRepository{
		databaseClient: databaseClient,
		logger:         logger,
	}
}

func (r *conversationRepository) Create(data domain.Conversation) (*domain.Conversation, error) {
	entity := ToConversationEntityGorm(data)

	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}
	if err := writer.Create(&entity).Error; err != nil {
		return nil, err
	}
	finalResponse := ToConversationDomain(entity)

	return &finalResponse, nil
}

func (r *conversationRepository) GetById(id uint64) (*domain.Conversation, error) {
	reader := r.databaseClient.ReadDB

	var entity ConversationEntityGorm
	err := reader.Model(&ConversationEntityGorm{}).
		Where("id = ?", id).
		First(&entity).Error
	if err != nil {
		return nil, err
	}
	finalResponse := ToConversationDomain(entity)

	return &finalResponse, nil
}
