package di

import (
	http3 "binar/internal/app/conversation/delivery/http"
	http4 "binar/internal/app/files/delivery/http"
	http2 "binar/internal/app/notifications/delivery/http"
	"binar/internal/app/users/delivery/http"
	"binar/internal/infra"
	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	http.NewUserHandler,
	http2.NewNotificationHandler,
	http3.NewConversationHandler,
	http4.NewFileHandler,
	ProvideHandler,
)

func ProvideHandler(
	userHandler *http.UserHandler,
	notificationHandler *http2.NotificationHandler,
	conversationHandler *http3.ConversationHandler,
	fileHandler *http4.FileHandler,
) infra.Handler {
	return infra.Handler{
		UserHandler:         *userHandler,
		NotificationHandler: *notificationHandler,
		ConversationHandler: *conversationHandler,
		FileHandler:         *fileHandler,
	}
}
