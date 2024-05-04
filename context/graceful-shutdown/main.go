package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	go func() {
		InfiniteForLoop(ctx, time.Second)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
		log.Println("Singint got")
		cancel()
		time.Sleep(time.Second) // this will give other goroutines that listen context some time to gracefully finish
		log.Println("Gracefully stoped")
	}
}

func InfiniteForLoop(ctx context.Context, interval time.Duration) {
loop:
	for {
		select {
		case <-ctx.Done():
			log.Println("Breaking loop")
			break loop
		case <-time.After(time.Second):
			log.Println("iterating...")
		}
	}
}
