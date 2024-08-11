package domain

import "time"

type Notification struct {
	Id        uint64
	UserId    uint64
	Message   string
	SentAt    time.Time
	Broadcast bool
}

type Job struct {
	Id          uint64
	Status      string
	QueuedAt    time.Time
	CompletedAt *time.Time
	Message     string
}

type NotificationRepository interface {
	Create(data Notification) (*Notification, error)
	GetAllByUserId(userId uint64) ([]Notification, error)
	CreateBroadcastNotifications(message string) error
	SaveJob(job Job) error
	GetJobByID(jobID string) (*Job, error)
	UpdateJobStatus(jobID string, status string, completedAt *time.Time) error
}

type NotificationService interface {
	Create(data Notification) (*Notification, string, error)
	GetAllByUserId(userId uint64) ([]Notification, string, error)
	BroadcastNotification(message string) (*Job, string, error)
	GetJobStatus(jobID string) (*Job, string, error)
}
