package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"randomaliens/client"
	"randomaliens/internal/repository"
	"randomaliens/internal/service"
	"syscall"
	"time"
)

func main() {
	db, err := repository.NewPostgresDb()
	if err != nil {
		log.Fatalf(err.Error())
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	client := client.NewClient(service)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	go func() {
		err := client.StartTransmission(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGINT)
	log.Println("Registered interrupt handler for client")

	switch <-osSignals {
	default:
		log.Printf("Handling interrupt singnal...")
		cancel()
		time.Sleep(time.Second)
		log.Println("Gracefuly stopping finished.")
	}
}
