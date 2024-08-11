package infra

import (
	http3 "binar/internal/app/conversation/delivery/http"
	domain3 "binar/internal/app/conversation/domain"
	http4 "binar/internal/app/files/delivery/http"
	domain4 "binar/internal/app/files/domain"
	http2 "binar/internal/app/notifications/delivery/http"
	domain2 "binar/internal/app/notifications/domain"
	queueService "binar/internal/app/queue/domain"
	"binar/internal/app/users/delivery/http"
	"binar/internal/app/users/domain"
	"binar/internal/infra/database"
	"binar/pkg/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Infra struct {
	Config       *config.Config
	Databases    *database.Databases
	RabbitMQ     *amqp.Connection
	Logger       *zap.Logger
	Service      Service
	Repository   Repository
	Handler      Handler
	QueueService queueService.QueueService
}

type Service struct {
	UserService         domain.UserService
	NotificationService domain2.NotificationService
	ConversationService domain3.ConversationService
	MessageService      domain3.MessageService
	FileService         domain4.FileService
}

type Repository struct {
	UserRepository         domain.UserRepository
	NotificationRepository domain2.NotificationRepository
	ConversationRepository domain3.ConversationRepository
	MessageRepository      domain3.MessageRepository
	FileRepository         domain4.FileRepository
}

type Handler struct {
	UserHandler         http.UserHandler
	NotificationHandler http2.NotificationHandler
	ConversationHandler http3.ConversationHandler
	FileHandler         http4.FileHandler
}
