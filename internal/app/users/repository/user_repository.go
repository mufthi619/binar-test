package repository

import (
	"binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"binar/internal/infra/gorm"
	"go.uber.org/zap"
	gorm2 "gorm.io/gorm"
	"time"
)

type userRepository struct {
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewUserRepository(databaseClient *database.Databases, logger *zap.Logger) domain.UserRepository {
	return &userRepository{
		databaseClient: databaseClient,
		logger:         logger,
	}
}

func (r *userRepository) Create(data domain.User) (*domain.User, error) {
	entity := ToUserEntityGorm(data)
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()

	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}
	if err := writer.Create(&entity).Error; err != nil {
		return nil, err
	}
	finalResponse := ToDomainUser(entity)

	return &finalResponse, nil
}

func (r *userRepository) GetById(id uint64) (*domain.User, error) {
	reader := r.databaseClient.ReadDB

	var resp UserEntityGorm
	err := reader.Scopes(gorm.FilterSoftDelete).First(&resp, id).Error
	if err != nil {
		return nil, err
	}
	finalResponse := ToDomainUser(resp)

	return &finalResponse, nil
}

func (r *userRepository) GetByIdIn(id []uint64) ([]domain.User, error) {
	reader := r.databaseClient.ReadDB

	var resp []UserEntityGorm
	query := reader.Scopes(gorm.FilterSoftDelete).Find(&resp, id)
	if query.Error != nil {
		return nil, query.Error
	}

	if len(resp) == 0 {
		return nil, gorm2.ErrRecordNotFound
	}

	finalResponse := ToDomainUsers(resp)
	return finalResponse, nil
}
