package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"randomaliens/server"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		grpcServer := server.NewServer()
		if err := grpcServer.Start(ctx, "8000"); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	select {
	case <-osSignals:
		log.Printf("Handling interrupt singnal...")
		time.Sleep(time.Second)
		cancel()
		log.Printf("Finished shutting server down.")
	}
}
