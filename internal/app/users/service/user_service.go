package service

import (
	"binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"binar/internal/infra/gorm"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	gorm2 "gorm.io/gorm"
)

type userService struct {
	userRepository domain.UserRepository
	databaseClient *database.Databases
	logger         *zap.Logger
}

func NewUserService(userRepo domain.UserRepository, databaseClient *database.Databases, logger *zap.Logger) domain.UserService {
	return &userService{
		userRepository: userRepo,
		databaseClient: databaseClient,
		logger:         logger,
	}
}

const (
	dbProblem    = "Failed ! There's some trouble in our system, please try again"
	successfully = "Successfully"
)

func (s *userService) Create(data domain.User) (*domain.User, string, error) {

	//Password Hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-1] | Failed on bcrypt.GenerateFromPassword, err -> %v", err)
		return nil, dbProblem, err
	}
	data.Password = string(hashedPassword)

	//Tx Section
	tx := s.databaseClient.WriteDB.Begin()
	if err = tx.Error; err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-2] | Failed on opentx, err -> %v", err)
		return nil, dbProblem, err
	}
	s.databaseClient.TxDb = tx
	defer tx.Rollback()

	//Create
	resp, err := s.userRepository.Create(data)
	if err != nil {
		if is, column := gorm.IsDuplicateError(err); is {
			s.logger.Sugar().Errorf("[Create][Flag-3] | Duplicate data, err -> %v", err)
			return nil, fmt.Sprintf("Failed ! %v already exists", column), err
		}
		s.logger.Sugar().Errorf("[Create][Flag-4] | Failed on userRepository.Create, err -> %v", err)
		return nil, dbProblem, err
	}

	//Commit
	err = tx.Commit().Error
	if err != nil {
		s.logger.Sugar().Errorf("[Create][Flag-5] | Failed on commit, err -> %v", err)
		return nil, dbProblem, err
	}

	//Final Response
	return resp, successfully, nil
}

func (s *userService) GetById(id uint64) (*domain.User, string, error) {
	resp, err := s.userRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm2.ErrRecordNotFound) {
			return resp, successfully, nil
		}
		s.logger.Sugar().Errorf("[GetById][Flag-1] | Failed on userRepository.GetById, err -> %v", err)
		return nil, dbProblem, err
	}

	return resp, successfully, nil
}
