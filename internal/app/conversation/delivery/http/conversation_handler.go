package http

import (
	payload2 "binar/internal/app/conversation/delivery/http/payload"
	"binar/internal/app/conversation/domain"
	"binar/internal/infra/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ConversationHandler struct {
	messageService      domain.MessageService
	conversationService domain.ConversationService
}

func NewConversationHandler(messageService domain.MessageService, conversationService domain.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		messageService:      messageService,
		conversationService: conversationService,
	}
}

func (h *ConversationHandler) CreateMessage(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var dataCreate payload2.CreateMessageRequest
	if err := ctx.Bind(&dataCreate); err != nil {
		log.Print(1, err)
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	if err := ctx.Validate(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest(err.Error())
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	messageDomain := domain.Message{
		ConversationId: dataCreate.ConversationId,
		SenderId:       dataCreate.SenderId,
		Content:        dataCreate.Content,
	}

	resp, msg, err := h.messageService.Create(messageDomain)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload2.MessageDetailResponse{
		Id:             resp.Id,
		ConversationId: resp.ConversationId,
		SenderId:       resp.SenderId,
		Content:        resp.Content,
		SentAt:         resp.SentAt,
	}
	responseFormatter.ReturnCreatedWithData(finalResponse, msg)

	return ctx.JSON(http.StatusCreated, responseFormatter)
}

func (h *ConversationHandler) GetMessagesByConversation(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var filter payload2.FindMessagesByConversation
	if err := ctx.Bind(&filter); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.messageService.GetAllByConversationId(filter.ConversationId)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable()
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	var finalResponse []payload2.MessageDetailResponse
	for _, val := range resp {
		finalResponse = append(finalResponse, payload2.MessageDetailResponse{
			Id:             val.Id,
			ConversationId: val.ConversationId,
			SenderId:       val.SenderId,
			Content:        val.Content,
			SentAt:         val.SentAt,
		})
	}
	responseFormatter.ReturnSuccessfullyWithData(finalResponse, msg)

	return ctx.JSON(http.StatusOK, responseFormatter)
}

func (h *ConversationHandler) CreateConversation(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var dataCreate payload2.CreateConversationRequest
	if err := ctx.Bind(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}
	if err := ctx.Validate(&dataCreate); err != nil {
		responseFormatter.ReturnBadRequest(err.Error())
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.conversationService.Create(dataCreate.Participants)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable(msg)
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload2.ConversationDetailResponse{
		Id:           resp.Id,
		Participants: resp.Participants,
		CreatedAt:    resp.CreatedAt,
	}
	responseFormatter.ReturnCreatedWithData(finalResponse, msg)

	return ctx.JSON(http.StatusCreated, responseFormatter)
}

func (h *ConversationHandler) GetConversationById(ctx echo.Context) error {
	responseFormatter := response.NewResponseFormatter()

	var filter payload2.FindConversation
	if err := ctx.Bind(&filter); err != nil {
		responseFormatter.ReturnBadRequest()
		return ctx.JSON(http.StatusBadRequest, responseFormatter)
	}

	resp, msg, err := h.conversationService.GetById(filter.Id)
	if err != nil {
		responseFormatter.ReturnInternalUnavailable()
		return ctx.JSON(http.StatusServiceUnavailable, responseFormatter)
	}

	finalResponse := payload2.ConversationDetailResponse{
		Id:           resp.Id,
		Participants: resp.Participants,
		CreatedAt:    resp.CreatedAt,
	}
	responseFormatter.ReturnSuccessfullyWithData(finalResponse, msg)

	return ctx.JSON(http.StatusOK, responseFormatter)
}
