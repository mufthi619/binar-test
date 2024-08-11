package repository

import (
	"binar/internal/app/users/domain"
	"time"
)

func ToUserEntityGorm(user domain.User) UserEntityGorm {
	return UserEntityGorm{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToDomainUser(entity UserEntityGorm) domain.User {
	var deletedAt *time.Time
	if entity.DeletedAt != nil {
		if !entity.DeletedAt.IsZero() {
			deletedAt = entity.DeletedAt
		}
	}
	return domain.User{
		ID:        entity.ID,
		Username:  entity.Username,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

func ToDomainUsers(entity []UserEntityGorm) (finalResponse []domain.User) {
	for _, val := range entity {
		finalResponse = append(finalResponse, ToDomainUser(val))
	}
	return
}
