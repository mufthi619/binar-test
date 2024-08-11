package seeder

import (
	"binar/internal/app/users/domain"
	"binar/internal/app/users/repository"
	"fmt"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserSeeder(db *gorm.DB, logger *zap.Logger) *UserSeeder {
	return &UserSeeder{
		db:     db,
		logger: logger,
	}
}

func (s *UserSeeder) Seed() error {
	users := []struct {
		username string
		email    string
		password string
	}{
		{"satoru_gojo", "gojo@jjk.com", "the_strongest_man_in_nowday"},
		{"yuji_itadori", "yuji@jjk.com", "mcgajelas"},
		{"megumi_fushiguro", "megumi@jjk.com", "bebandahiniorang"},
		{"maki_zenin", "maki@jjk.com", "tibatibaggdahiniorang"},
	}

	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %v", err)
		}

		userDomain := domain.User{
			Username:  user.username,
			Email:     user.email,
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		userEntity := repository.ToUserEntityGorm(userDomain)
		if err := s.db.Create(&userEntity).Error; err != nil {
			return fmt.Errorf("error seeding user %s: %v", user.username, err)
		}
	}

	s.logger.Info("User seeding completed successfully")
	return nil
}
