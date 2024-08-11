package service

import (
	"binar/internal/app/notifications/domain"
	domain3 "binar/internal/app/queue/domain"
	domain2 "binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"errors"
	"go.uber.org/zap"
	gorm2 "gorm.io/gorm"
	"time"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	userRepository         domain2.UserRepository
	databaseClient         *database.Databases
	queueService           domain3.QueueService
	logger                 *zap.Logger
}

func NewNotificationService(
	notificationRepo domain.NotificationRepository,
	userRepo domain2.UserRepository,
	databaseClient *database.Databases,
	queueService domain3.QueueService,
	logger *zap.Logger,
) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepo,
		userRepository:         userRepo,
		databaseClient:         databaseClient,
		logger:                 logger,
		queueService:           queueService,
	}
}

const (
	dbProblem    = "Failed ! There's some trouble in our system, please try again"
	successfully = "Successfully"
)

func (s *notificationService) Create(data domain.Notification) (*domain.Notification, string, error) {
	//Validate UserId
	_, err := s.userRepository.GetByIdIn([]uint64{data.UserId})
	if err != nil {
		if errors.Is(err, gorm2.ErrRecordNotFound) {
			return nil, "Failed ! Please make sure user is registered", err
		}
		s.logger.Sugar().Errorf("[Create][Flag-1] | Failed on userRepository.GetByIdIn, err -> %v", err)
		return nil, dbProblem, err
	}

	//Tx Section
	tx := s.databaseClient.WriteDB.Begin()
	if err := tx.Error; err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-1] | Failed on opentx, err -> %v", err)
		return nil, dbProblem, err
	}
	s.databaseClient.TxDb = tx
	defer tx.Rollback()

	//Create
	resp, err := s.notificationRepository.Create(data)
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-2] | Failed on notificationRepository.Create, err -> %v", err)
		return nil, dbProblem, tx.Error
	}

	//Commit
	err = tx.Commit().Error
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-3] | Failed on commit, err -> %v", err)
		return nil, dbProblem, err
	}

	//Final Response
	return resp, successfully, nil
}

func (s *notificationService) GetAllByUserId(userId uint64) ([]domain.Notification, string, error) {
	resp, err := s.notificationRepository.GetAllByUserId(userId)
	if err != nil {
		if errors.Is(err, gorm2.ErrRecordNotFound) {
			return resp, successfully, nil
		}
		s.logger.Sugar().Errorf("[GetById][Flag-1] | Failed on notificationRepository.GetAllByUserId, err -> %v", err)
		return nil, dbProblem, err
	}

	return resp, successfully, nil
}

func (s *notificationService) BroadcastNotification(message string) (*domain.Job, string, error) {
	// Create a new job
	job := domain.Job{
		Status:   "queued",
		QueuedAt: time.Now(),
		Message:  message,
	}

	// Save the job to the database
	err := s.notificationRepository.SaveJob(job)
	if err != nil {
		s.logger.Sugar().Errorf("[BroadcastNotification][Flag-1] | Failed to save job, err -> %v", err)
		return nil, dbProblem, err
	}

	// Publish the message to the queue
	queueMessage := domain3.QueueMessage{
		Type:    "broadcast_notification",
		Payload: []byte(message),
	}
	err = s.queueService.PublishMessage(queueMessage)
	if err != nil {
		s.logger.Sugar().Errorf("[BroadcastNotification][Flag-2] | Failed to publish message to queue, err -> %v", err)
		return nil, dbProblem, err
	}

	return &job, successfully, nil
}

func (s *notificationService) GetJobStatus(jobID string) (*domain.Job, string, error) {
	job, err := s.notificationRepository.GetJobByID(jobID)
	if err != nil {
		if errors.Is(err, gorm2.ErrRecordNotFound) {
			return nil, "Job not found", err
		}
		s.logger.Sugar().Errorf("[GetJobStatus][Flag-1] | Failed to get job, err -> %v", err)
		return nil, dbProblem, err
	}

	return job, successfully, nil
}
