package di

import (
	domain5 "binar/internal/app/article/category/domain"
	repository4 "binar/internal/app/article/category/repository"
	domain3 "binar/internal/app/conversation/domain"
	"binar/internal/app/conversation/repository/conversation"
	"binar/internal/app/conversation/repository/message"
	domain4 "binar/internal/app/files/domain"
	repository3 "binar/internal/app/files/repository"
	domain2 "binar/internal/app/notifications/domain"
	repository2 "binar/internal/app/notifications/repository"
	"binar/internal/app/users/domain"
	"binar/internal/app/users/repository"
	"binar/internal/infra"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository2.NewNotificationRepository,
	conversation.NewConversationRepository,
	message.NewMessageRepository,
	repository3.NewFileRepository,
	repository4.NewCategoryRepository,
	ProvideRepository,
)

func ProvideRepository(
	userRepo domain.UserRepository,
	notificationRepo domain2.NotificationRepository,
	conversationRepository domain3.ConversationRepository,
	messageRepository domain3.MessageRepository,
	fileRepository domain4.FileRepository,
	categoryRepository domain5.CategoryRepository,
) infra.Repository {
	return infra.Repository{
		UserRepository:         userRepo,
		NotificationRepository: notificationRepo,
		ConversationRepository: conversationRepository,
		MessageRepository:      messageRepository,
		FileRepository:         fileRepository,
		CategoryRepository:     categoryRepository,
	}
}
