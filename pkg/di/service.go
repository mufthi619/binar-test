package di

import (
	domain3 "binar/internal/app/conversation/domain"
	service3 "binar/internal/app/conversation/service"
	domain4 "binar/internal/app/files/domain"
	service4 "binar/internal/app/files/service"
	domain2 "binar/internal/app/notifications/domain"
	service2 "binar/internal/app/notifications/service"
	queueService "binar/internal/app/queue/service"
	"binar/internal/app/users/domain"
	"binar/internal/app/users/service"
	"binar/internal/infra"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	service.NewUserService,
	service2.NewNotificationService,
	service3.NewConversationService,
	service3.NewMessageService,
	service4.NewFileService,
	ProvideService,
	queueService.NewQueueService,
)

func ProvideService(
	userService domain.UserService,
	notificationService domain2.NotificationService,
	conversationService domain3.ConversationService,
	messageService domain3.MessageService,
	fileService domain4.FileService,
) infra.Service {
	return infra.Service{
		UserService:         userService,
		NotificationService: notificationService,
		ConversationService: conversationService,
		MessageService:      messageService,
		FileService:         fileService,
	}
}
