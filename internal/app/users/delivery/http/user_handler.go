package http

import (
	"binar/internal/app/users/delivery/http/payload"
	"binar/internal/app/users/domain"
	"binar/internal/infra/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var dataCreate payload.CreateUserRequest
	if err := ctx.Bind(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}
	if err := ctx.Validate(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest(err.Error())
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	userDomain := domain.User{
		Username: dataCreate.Username,
		Email:    dataCreate.Email,
		Password: dataCreate.Password,
	}

	resp, msg, err := h.userService.Create(userDomain)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload.UserDetailResponse{
		Id:        resp.ID,
		Username:  resp.Username,
		Email:     resp.Email,
		CreatedAt: resp.CreatedAt,
	}
	responseFormatter.ReturnCreatedWithData(finalResponse, msg)

	return ctx.JSON(http.StatusCreated, responseFormatter)
}

func (h *UserHandler) FindUserById(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var filter payload.FindUser
	if err := ctx.Bind(&filter); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.userService.GetById(filter.Id)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable()
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}
	responseFormatter.ReturnSuccessfullyWithData(resp, msg)

	return ctx.JSON(http.StatusOK, responseFormatter)
}
