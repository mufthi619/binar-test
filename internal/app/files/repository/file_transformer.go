package repository

import (
	"binar/internal/app/files/domain"
)

func ToFileEntityGorm(data domain.File) FileEntityGorm {
	return FileEntityGorm{
		Id:         data.Id,
		UserId:     data.UserId,
		FileUrl:    data.FileUrl,
		UploadedAt: data.UploadedAt,
	}
}

func ToFileDomain(entity FileEntityGorm) domain.File {
	return domain.File{
		Id:         entity.Id,
		UserId:     entity.UserId,
		FileUrl:    entity.FileUrl,
		UploadedAt: entity.UploadedAt,
	}
}

func ToFilesDomain(entities []FileEntityGorm) (finalResponse []domain.File) {
	for _, entity := range entities {
		finalResponse = append(finalResponse, ToFileDomain(entity))
	}
	return
}
