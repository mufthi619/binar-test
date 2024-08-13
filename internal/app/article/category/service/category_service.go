package service

import (
	"binar/internal/app/article/category/domain"
	"binar/internal/infra/database"
	"go.uber.org/zap"
)

type categoryService struct {
	categoryRepository domain.CategoryRepository
	databaseClient     *database.Databases
	logger             *zap.Logger
}

func NewCategoryService(categoryRepo domain.CategoryRepository, databaseClient *database.Databases, logger *zap.Logger) domain.CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
		databaseClient:     databaseClient,
		logger:             logger,
	}
}

const (
	dbProblem    = "Failed ! There's some trouble in our system, please try again"
	successfully = "Successfully"
)

func (s *categoryService) GetAll() ([]domain.Category, string, error) {
	resp, err := s.categoryRepository.GetAll()
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-1] | Failed on categoryRepository.GetAll(), err -> %v", err)
		return nil, dbProblem, err
	}

	return resp, successfully, err
}
