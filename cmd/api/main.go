package main

import (
	"binar/internal/infra/server"
	"binar/pkg/config"
	"binar/pkg/di"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig("config/app.yaml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	app, err := di.InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	srv := server.NewServer(app)
	go func() {
		if err := srv.Start(); err != nil {
			app.Logger.Sugar().Infof("Failed to start server : %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	app.Logger.Info("Server is shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		app.Logger.Sugar().Errorf("Server forced to shutdown : %v", err)
	}

	app.Logger.Info("Server exited properly")
}
