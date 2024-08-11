package payload

import (
	"time"
)

type UserDetailResponse struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
