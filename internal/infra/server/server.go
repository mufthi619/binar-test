package server

import (
	"binar/internal/infra"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"path/filepath"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type Server struct {
	app *infra.Infra
	e   *echo.Echo
}

func NewServer(app *infra.Infra) *Server {
	e := echo.New()
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	uploadsDir := filepath.Join(".", app.Config.AppConfig.PublicDir)
	e.Static("/uploads", uploadsDir)

	server := &Server{
		app: app,
		e:   e,
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	api := s.e.Group("/api")

	//User
	api.POST("/users", s.app.Handler.UserHandler.CreateUser)
	api.GET("/users/:id", s.app.Handler.UserHandler.FindUserById)

	//Notification
	api.POST("/notifications", s.app.Handler.NotificationHandler.CreateNotification)
	api.GET("/notifications/:user_id", s.app.Handler.NotificationHandler.FindNotification)
	api.POST("/notifications/broadcast", s.app.Handler.NotificationHandler.BroadcastNotification)

	//Conversation
	api.POST("/conversations", s.app.Handler.ConversationHandler.CreateConversation)
	api.GET("/conversations/:id", s.app.Handler.ConversationHandler.GetConversationById)
	api.POST("/conversations/:conversation_id/messages", s.app.Handler.ConversationHandler.CreateMessage)
	api.GET("/conversations/:conversation_id/messages", s.app.Handler.ConversationHandler.GetMessagesByConversation)

	//Files
	api.POST("/files/upload", s.app.Handler.FileHandler.UploadFile)
	api.GET("/files/:id", s.app.Handler.FileHandler.GetFileById)

	//Job
	api.GET("/jobs/:id", s.app.Handler.NotificationHandler.GetJobStatus)

	//Category
	api.GET("/category", s.app.Handler.CategoryHandler.GetAll)

}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.app.Config.AppConfig.Port)
	return s.e.Start(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
