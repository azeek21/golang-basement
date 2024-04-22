package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"server"
	"server/pkg/handlers"
	"server/pkg/repository"
	"server/pkg/service"
	"server/types"
)

func main() {
	if err := types.LoadGlobalConfig(); err != nil {
		log.Fatal(err.Error())
		return
	}

	elastic, err := repository.NewElasticClient(&types.GLOABAL_CONFIG.ElasticEnv)
	if err != nil {
		log.Fatal(err)
		return
	}

	postgres, err := repository.NewPostgresDb(types.GLOABAL_CONFIG.PostgresDbConfig)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	repo := repository.NewRepository(postgres, elastic)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service)
	server := new(server.Server)

	go func() {
		if err := server.Run("8888", handler.InitRoutes()); err != nil {
			log.Fatal(err.Error())
			return
		}
	}()

	log.Println("Restaurants app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("TodoApp Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured during shutting down the server: %s", err.Error())
	}

	if db, err := postgres.DB(); err != nil {
		log.Fatalf("error getting underlying db from gorm %s", err.Error())
	} else {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing postgres connection %s", err.Error())
		}
	}
}
