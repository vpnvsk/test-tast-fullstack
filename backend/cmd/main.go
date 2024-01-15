package main

import (
	"context"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/internal/server"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/pkg/handler"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/pkg/repository"
	"os"
	"os/signal"
	"syscall"
)

// @title Test-task Backend
// @version 1.0
// @description API Server for test-task

// @host localhost:8000
// @BasePath /
func main() {

	db := repository.NewInMemoryDB()
	repo := repository.NewRepository(db)
	handl := handler.NewHandler(repo)
	srv := new(server.Server)
	go func() {
		if err := srv.Run("8000", handl.InitRoutes()); err != nil {
			panic("error while starting server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.ShutDown(context.Background()); err != nil {

	}
}
