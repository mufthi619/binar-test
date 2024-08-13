package http

import (
	"binar/internal/app/article/category/delivery/http/payload"
	"binar/internal/app/article/category/domain"
	"binar/internal/infra/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryHandler struct {
	categoryHandler domain.CategoryService
}

func NewCategoryHandler(service domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryHandler: service}
}

func (h *CategoryHandler) GetAll(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	resp, msg, err := h.categoryHandler.GetAll()
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	var finalResponse payload.GetCategoryResponse
	var finalResponseData []payload.ArticleGroup
	for _, val := range resp {
		finalResponseData = append(finalResponseData, payload.ArticleGroup{
			Id:   val.Id,
			Name: val.Name,
		})
	}
	finalResponse.ArticleCategory = finalResponseData

	responseFormatter.ReturnSuccessfullyWithData(finalResponse, msg)
	return ctx.JSON(http.StatusOK, responseFormatter)
}
