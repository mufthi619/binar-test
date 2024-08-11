package service

import (
	"binar/internal/app/conversation/domain"
	domain2 "binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type conversationService struct {
	conversationRepository domain.ConversationRepository
	userRepository         domain2.UserRepository
	databaseClient         *database.Databases
	logger                 *zap.Logger
}

func NewConversationService(
	conversationRepo domain.ConversationRepository,
	userRepo domain2.UserRepository,
	databaseClient *database.Databases,
	logger *zap.Logger,
) domain.ConversationService {
	return &conversationService{
		conversationRepository: conversationRepo,
		userRepository:         userRepo,
		databaseClient:         databaseClient,
		logger:                 logger,
	}
}

func (s *conversationService) Create(participants []uint64) (*domain.Conversation, string, error) {
	// Validate participants
	users, err := s.userRepository.GetByIdIn(participants)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "Failed ! Please make sure all participants are registered", err
		}
		s.logger.Sugar().Errorf("[Create][Flag-1] | Failed on userRepository.GetByIdIn, err -> %v", err)
		return nil, dbProblem, err
	}
	if len(users) != len(participants) {
		return nil, "Failed ! Some participants are not registered", errors.New("invalid participants")
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
	conversation := domain.Conversation{Participants: participants}
	resp, err := s.conversationRepository.Create(conversation)
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-3] | Failed on conversationRepository.Create, err -> %v", err)
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

func (s *conversationService) GetById(id uint64) (*domain.Conversation, string, error) {
	resp, err := s.conversationRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "Conversation not found", err
		}
		s.logger.Sugar().Errorf("[GetById][Flag-1] | Failed on conversationRepository.GetById, err -> %v", err)
		return nil, dbProblem, err
	}

	return resp, successfully, nil
}
