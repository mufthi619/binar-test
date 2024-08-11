package http

import (
	"binar/internal/app/files/delivery/http/payload"
	"binar/internal/app/files/domain"
	"binar/internal/infra/response"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type FileHandler struct {
	fileService domain.FileService
}

func NewFileHandler(service domain.FileService) *FileHandler {
	return &FileHandler{fileService: service}
}

func (h *FileHandler) UploadFile(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	file, err := ctx.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			responseFormatter.ReturnBadRequest("No file uploaded")
		} else {
			responseFormatter.ReturnBadRequest("Failed to get file from request")
		}
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}
	if file.Size > 10*1024*1024 {
		responseFormatter.ReturnBadRequest("File too large")
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	var dataUpload payload.UploadFile
	if err := ctx.Bind(&dataUpload); err != nil {
		responseFormatter.ReturnBadRequest("Invalid form data")
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	if err := ctx.Validate(&dataUpload); err != nil {
		responseFormatter.ReturnBadRequest(err.Error())
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	src, err := file.Open()
	if err != nil {
		responseFormatter.ReturnBadRequest("Failed to open file")
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}
	defer src.Close()

	fileData, err := ioutil.ReadAll(src)
	if err != nil {
		responseFormatter.ReturnBadRequest("Failed to read file")
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.fileService.Upload(dataUpload.UserId, fileData, file.Filename)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload.FileDetailResponse{
		Id:         resp.Id,
		UserId:     resp.UserId,
		FileUrl:    resp.FileUrl,
		UploadedAt: resp.UploadedAt,
	}
	responseFormatter.ReturnCreatedWithData(finalResponse, msg)

	return ctx.JSON(http.StatusCreated, responseFormatter)
}

func (h *FileHandler) GetFileById(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var filter payload.FindFile
	if err := ctx.Bind(&filter); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.fileService.GetById(filter.Id)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload.FileDetailResponse{
		Id:         resp.Id,
		UserId:     resp.UserId,
		FileUrl:    resp.FileUrl,
		UploadedAt: resp.UploadedAt,
	}
	responseFormatter.ReturnSuccessfullyWithData(finalResponse, msg)

	return ctx.JSON(http.StatusOK, responseFormatter)
}
