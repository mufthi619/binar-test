package domain

import "time"

type File struct {
	Id         uint64
	UserId     uint64
	FileUrl    string
	UploadedAt time.Time
}

type FileRepository interface {
	GetById(id uint64) (*File, error)
	SaveFileInfo(userId uint64, fileUrl string) (*File, error)
}

type FileService interface {
	Upload(userId uint64, fileData []byte, originalFileName string) (*File, string, error)
	GetById(id uint64) (*File, string, error)
}
