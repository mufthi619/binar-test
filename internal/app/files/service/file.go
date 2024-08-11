package service

import (
	"binar/internal/app/files/domain"
	domain2 "binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"binar/pkg/config"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
	"path/filepath"
	"time"
)

type fileService struct {
	fileRepository domain.FileRepository
	userRepository domain2.UserRepository
	databaseClient *database.Databases
	appConfig      *config.AppConfig
	logger         *zap.Logger
}

func NewFileService(
	fileRepo domain.FileRepository,
	userRepo domain2.UserRepository,
	databaseClient *database.Databases,
	appConfig *config.AppConfig,
	logger *zap.Logger,
) domain.FileService {
	return &fileService{
		fileRepository: fileRepo,
		userRepository: userRepo,
		databaseClient: databaseClient,
		appConfig:      appConfig,
		logger:         logger,
	}
}

const (
	dbProblem    = "Failed ! There's some trouble in our system, please try again"
	successfully = "Successfully"
)

func (s *fileService) Upload(userId uint64, fileData []byte, originalFileName string) (*domain.File, string, error) {
	// Validate UserId
	_, err := s.userRepository.GetByIdIn([]uint64{userId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "Failed ! Please make sure user is registered", err
		}
		s.logger.Sugar().Errorf("[Upload][Flag-1] | Failed on userRepository.GetByIdIn, err -> %v", err)
		return nil, dbProblem, err
	}

	// Tx Section
	tx := s.databaseClient.WriteDB.Begin()
	if err := tx.Error; err != nil {
		s.logger.Sugar().Errorf("[Upload][Flag-1] | Failed on opentx, err -> %v", err)
		return nil, dbProblem, err
	}
	s.databaseClient.TxDb = tx
	defer tx.Rollback()

	//File Upload
	fileExt := filepath.Ext(originalFileName)
	fileName := fmt.Sprintf("%d_%d%s", userId, time.Now().UnixNano(), fileExt)
	filePath := filepath.Join(s.appConfig.PublicDir, fileName)
	err = ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		s.logger.Sugar().Errorf("[Upload][Flag-2] | Failed to save file locally, err -> %v", err)
		return nil, dbProblem, err
	}

	//Repo Save File Info
	resp, err := s.fileRepository.SaveFileInfo(userId, fileName)
	if err != nil {
		s.logger.Sugar().Errorf("[Upload][Flag-3] | Failed on fileRepository.Upload, err -> %v", err)
		return nil, dbProblem, tx.Error
	}
	if s.appConfig.Port != 0 {
		resp.FileUrl = fmt.Sprintf("%s:%v/%s/%s", s.appConfig.URL, s.appConfig.Port, s.appConfig.PublicDir, fileName)
	} else {
		resp.FileUrl = fmt.Sprintf("%s/%s/%s", s.appConfig.URL, s.appConfig.PublicDir, fileName)
	}

	// Commit
	err = tx.Commit().Error
	if err != nil {
		s.logger.Sugar().Errorf("[Upload][Flag-4] | Failed on commit, err -> %v", err)
		return nil, dbProblem, err
	}

	// Final Response
	return resp, successfully, nil
}

func (s *fileService) GetById(id uint64) (*domain.File, string, error) {
	resp, err := s.fileRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "File not found", err
		}
		s.logger.Sugar().Errorf("[GetById][Flag-1] | Failed on fileRepository.GetById, err -> %v", err)
		return nil, dbProblem, err
	}

	return resp, successfully, nil
}
