package payload

type CreateNotification struct {
	UserId  uint64 `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type FindNotification struct {
	UserId uint64 `json:"user_id" query:"user_id" param:"user_id"`
}

type BroadcastNotificationRequest struct {
	Message string `json:"message" validate:"required"`
}
