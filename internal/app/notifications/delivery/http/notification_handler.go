package http

import (
	"binar/internal/app/notifications/delivery/http/payload"
	"binar/internal/app/notifications/domain"
	"binar/internal/infra/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NotificationHandler struct {
	notificationService domain.NotificationService
}

func NewNotificationHandler(service domain.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: service}
}

func (h *NotificationHandler) CreateNotification(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var dataCreate payload.CreateNotification
	if err := ctx.Bind(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}
	if err := ctx.Validate(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest(err.Error())
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	notificationDomain := domain.Notification{
		UserId:  dataCreate.UserId,
		Message: dataCreate.Message,
	}

	resp, msg, err := h.notificationService.Create(notificationDomain)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload.NotificationDetailResponse{
		Id:      resp.Id,
		UserId:  resp.UserId,
		Message: resp.Message,
		SentAt:  resp.SentAt,
	}
	responseFormatter.ReturnCreatedWithData(finalResponse, msg)

	return ctx.JSON(http.StatusCreated, responseFormatter)
}

func (h *NotificationHandler) FindNotification(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var filter payload.FindNotification
	if err := ctx.Bind(&filter); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.notificationService.GetAllByUserId(filter.UserId)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable()
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}
	var finalResponse payload.NotificationListResponse
	for _, val := range resp {
		finalResponse = append(finalResponse, payload.NotificationDetailResponse{
			Id:      val.Id,
			UserId:  val.UserId,
			Message: val.Message,
			SentAt:  val.SentAt,
		})
	}
	responseFormatter.ReturnSuccessfullyWithData(finalResponse, msg)

	return ctx.JSON(http.StatusOK, responseFormatter)
}

func (h *NotificationHandler) BroadcastNotification(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var dataCreate payload.BroadcastNotificationRequest
	if err := ctx.Bind(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}
	if err := ctx.Validate(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest(err.Error())
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	job, msg, err := h.notificationService.BroadcastNotification(dataCreate.Message)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload.BroadcastNotificationResponse{
		JobID:    job.Id,
		Status:   job.Status,
		QueuedAt: job.QueuedAt,
	}
	responseFormatter.ReturnCreatedWithData(finalResponse, msg)

	return ctx.JSON(http.StatusAccepted, responseFormatter)
}

func (h *NotificationHandler) GetJobStatus(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	jobID := ctx.Param("id")
	if jobID == "" {
		responseFormatter.ReturnBadRequest("Job ID is required")
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	job, msg, err := h.notificationService.GetJobStatus(jobID)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload.JobStatusResponse{
		ID:          job.Id,
		Status:      job.Status,
		QueuedAt:    job.QueuedAt,
		CompletedAt: job.CompletedAt,
		Message:     job.Message,
	}
	responseFormatter.ReturnSuccessfullyWithData(finalResponse, msg)

	return ctx.JSON(http.StatusOK, responseFormatter)
}
