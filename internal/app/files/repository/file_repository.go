package repository

import (
	"binar/internal/app/files/domain"
	"binar/internal/infra/database"
	"go.uber.org/zap"
	"time"
)

type fileRepository struct {
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewFileRepository(databaseClient *database.Databases, logger *zap.Logger) domain.FileRepository {
	return &fileRepository{
		databaseClient: databaseClient,
		logger:         logger,
	}
}

func (r *fileRepository) GetById(id uint64) (*domain.File, error) {
	reader := r.databaseClient.ReadDB

	var entity FileEntityGorm
	err := reader.First(&entity, id).Error
	if err != nil {
		return nil, err
	}

	file := ToFileDomain(entity)
	return &file, nil
}

func (r *fileRepository) SaveFileInfo(userId uint64, fileUrl string) (*domain.File, error) {
	entity := FileEntityGorm{
		UserId:     userId,
		FileUrl:    fileUrl,
		UploadedAt: time.Now(),
	}

	writer := r.databaseClient.WriteDB
	if tx := r.databaseClient.TxDb; tx != nil {
		if err := tx.Error; err == nil {
			writer = tx
		}
	}
	if err := writer.Create(&entity).Error; err != nil {
		return nil, err
	}

	file := ToFileDomain(entity)
	return &file, nil
}
