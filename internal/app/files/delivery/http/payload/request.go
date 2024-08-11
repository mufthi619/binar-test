package payload

import "mime/multipart"

type UploadFile struct {
	UserId uint64                `form:"user_id" validate:"required"`
	File   *multipart.FileHeader `form:"file"`
}

type FindFile struct {
	Id uint64 `json:"id" query:"id" param:"id"`
}

type FindFilesByUser struct {
	UserId uint64 `json:"user_id" query:"user_id" param:"user_id"`
}
