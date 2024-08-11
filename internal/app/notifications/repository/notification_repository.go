package repository

import (
	"binar/internal/app/notifications/domain"
	"binar/internal/infra/database"
	"go.uber.org/zap"
	"time"
)

type notificationRepository struct {
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewNotificationRepository(databaseClient *database.Databases, logger *zap.Logger) domain.NotificationRepository {
	return &notificationRepository{
		databaseClient: databaseClient,
		logger:         logger,
	}
}

func (r *notificationRepository) Create(data domain.Notification) (*domain.Notification, error) {
	entity := ToNotificationEntityGorm(data)

	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}
	if err := writer.Create(&entity).Error; err != nil {
		return nil, err
	}
	finalResponse := ToNotificationDomain(entity)

	return &finalResponse, nil
}

func (r *notificationRepository) GetAllByUserId(userId uint64) ([]domain.Notification, error) {
	reader := r.databaseClient.ReadDB

	var resp []NotificationEntityGorm
	err := reader.Model(&NotificationEntityGorm{}).
		Where("user_id = ?", userId).
		Scan(&resp).Error
	if err != nil {
		return nil, err
	}
	finalResponse := ToNotificationsDomain(resp)

	return finalResponse, nil
}

func (r *notificationRepository) CreateBroadcastNotifications(message string) error {
	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}

	err := writer.Exec(`
        INSERT INTO notifications (user_id, message, sent_at, broadcast)
        SELECT id, ?, NOW(), true FROM users
    `, message).Error

	return err
}

func (r *notificationRepository) SaveJob(job domain.Job) error {
	entity := ToJobEntityGorm(job)

	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}
	return writer.Create(&entity).Error
}

func (r *notificationRepository) GetJobByID(jobID string) (*domain.Job, error) {
	reader := r.databaseClient.ReadDB

	var jobEntity JobEntityGorm
	err := reader.Where("id = ?", jobID).First(&jobEntity).Error
	if err != nil {
		return nil, err
	}

	job := ToJobDomain(jobEntity)
	return &job, nil
}

func (r *notificationRepository) UpdateJobStatus(jobID string, status string, completedAt *time.Time) error {
	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}

	updates := map[string]interface{}{
		"status": status,
	}
	if completedAt != nil {
		updates["completed_at"] = completedAt
	}

	return writer.Model(&JobEntityGorm{}).Where("id = ?", jobID).Updates(updates).Error
}
