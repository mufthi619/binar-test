package payload

import "time"

type NotificationDetailResponse struct {
	Id      uint64    `json:"id"`
	UserId  uint64    `json:"user_id"`
	Message string    `json:"message"`
	SentAt  time.Time `json:"sent_at"`
}

type NotificationListResponse []NotificationDetailResponse

type BroadcastNotificationResponse struct {
	JobID    uint64    `json:"job_id"`
	Status   string    `json:"status"`
	QueuedAt time.Time `json:"queued_at"`
}

type JobStatusResponse struct {
	ID          uint64     `json:"id"`
	Status      string     `json:"status"`
	QueuedAt    time.Time  `json:"queued_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Message     string     `json:"message"`
}
