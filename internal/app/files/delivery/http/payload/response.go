package payload

import "time"

type FileDetailResponse struct {
	Id         uint64    `json:"id"`
	UserId     uint64    `json:"user_id"`
	FileUrl    string    `json:"file_url"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type FileListResponse []FileDetailResponse
