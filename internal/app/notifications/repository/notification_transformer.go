package repository

import (
	"binar/internal/app/notifications/domain"
)

func ToNotificationEntityGorm(data domain.Notification) NotificationEntityGorm {
	return NotificationEntityGorm{
		Id:      data.Id,
		UserId:  data.UserId,
		Message: data.Message,
		SentAt:  data.SentAt,
	}
}

func ToNotificationDomain(entity NotificationEntityGorm) domain.Notification {
	return domain.Notification{
		Id:      entity.Id,
		UserId:  entity.UserId,
		Message: entity.Message,
		SentAt:  entity.SentAt,
	}
}

func ToNotificationsDomain(entity []NotificationEntityGorm) (finalResponse []domain.Notification) {
	for _, val := range entity {
		resp := domain.Notification{
			Id:      val.Id,
			UserId:  val.UserId,
			Message: val.Message,
			SentAt:  val.SentAt,
		}
		finalResponse = append(finalResponse, resp)
	}
	return
}
