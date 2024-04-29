package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"candies/server"
	"candies/server/packages/handler"
	"candies/server/packages/repository"
	"candies/server/packages/service"
)

func main() {
	server := new(server.Server)
	repo := repository.NewRepository()
	service := service.NewService(*repo)
	handler := handler.NewHandler(*service)

	go func() {
		if err := server.Run("8082", handler.InitRoutes()); err != nil {
			log.Fatal(err.Error())
			return
		}
	}()

	log.Println("Candy server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
