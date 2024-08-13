package repository

import (
	"binar/internal/app/article/category/domain"
	"binar/internal/infra/database"
	"binar/internal/infra/gorm"
	"go.uber.org/zap"
)

type categoryRepository struct {
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewCategoryRepository(databaseClient *database.Databases, logger *zap.Logger) domain.CategoryRepository {
	return &categoryRepository{
		databaseClient: databaseClient,
		logger:         logger,
	}
}

func (r *categoryRepository) GetAll() ([]domain.Category, error) {
	reader := r.databaseClient.ReadDB

	var resp []CategoryEntityGorm
	err := reader.Table("categories").Scopes(gorm.FilterSoftDelete).Scan(&resp).Error
	if err != nil {
		return nil, err
	}
	finalResponse := ToCategoriesDomain(resp)

	return finalResponse, nil
}
