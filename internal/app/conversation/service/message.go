package service

import (
	"binar/internal/app/conversation/domain"
	domain2 "binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"binar/pkg/utils"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type messageService struct {
	messageRepository      domain.MessageRepository
	conversationRepository domain.ConversationRepository
	userRepository         domain2.UserRepository
	databaseClient         *database.Databases
	logger                 *zap.Logger
}

func NewMessageService(
	messageRepo domain.MessageRepository,
	conversationRepo domain.ConversationRepository,
	userRepo domain2.UserRepository,
	databaseClient *database.Databases,
	logger *zap.Logger,
) domain.MessageService {
	return &messageService{
		messageRepository:      messageRepo,
		conversationRepository: conversationRepo,
		userRepository:         userRepo,
		databaseClient:         databaseClient,
		logger:                 logger,
	}
}

func (s *messageService) Create(data domain.Message) (*domain.Message, string, error) {
	// Validate ConversationId and SenderId
	conversation, err := s.conversationRepository.GetById(data.ConversationId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "Failed ! Conversation not found", err
		}
		s.logger.Sugar().Errorf("[Create][Flag-1] | Failed on conversationRepository.GetById, err -> %v", err)
		return nil, dbProblem, err
	}
	if !utils.ContainsNumber(conversation.Participants, data.SenderId) {
		return nil, "Failed ! Sender is not a participant in this conversation", errors.New("unauthorized sender")
	}

	// Tx Section
	tx := s.databaseClient.WriteDB.Begin()
	if err := tx.Error; err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-2] | Failed on opentx, err -> %v", err)
		return nil, dbProblem, err
	}
	s.databaseClient.TxDb = tx
	defer tx.Rollback()

	// Create
	resp, err := s.messageRepository.Create(data)
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-3] | Failed on messageRepository.Create, err -> %v", err)
		return nil, dbProblem, tx.Error
	}

	// Commit
	err = tx.Commit().Error
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-4] | Failed on commit, err -> %v", err)
		return nil, dbProblem, err
	}

	// Final Response
	return resp, successfully, nil
}

func (s *messageService) GetAllByConversationId(conversationId uint64) ([]domain.Message, string, error) {
	resp, err := s.messageRepository.GetAllByConversationId(conversationId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, successfully, nil
		}
		s.logger.Sugar().Errorf("[GetAllByConversationId][Flag-1] | Failed on messageRepository.GetAllByConversationId, err -> %v", err)
		return nil, dbProblem, err
	}

	return resp, successfully, nil
}
