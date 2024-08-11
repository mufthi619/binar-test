package repository

import "binar/internal/app/notifications/domain"

func ToJobEntityGorm(data domain.Job) JobEntityGorm {
	return JobEntityGorm{
		Id:          data.Id,
		Status:      data.Status,
		QueuedAt:    data.QueuedAt,
		CompletedAt: data.CompletedAt,
		Message:     data.Message,
	}
}

func ToJobDomain(entity JobEntityGorm) domain.Job {
	return domain.Job{
		Id:          entity.Id,
		Status:      entity.Status,
		QueuedAt:    entity.QueuedAt,
		CompletedAt: entity.CompletedAt,
		Message:     entity.Message,
	}
}
