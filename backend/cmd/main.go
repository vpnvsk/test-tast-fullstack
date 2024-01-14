package main

import (
	"backend/internal/server"
	"backend/pkg/handler"
	"backend/pkg/repository"
	"context"
	"os"
	"os/signal"
	"syscall"
)

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
