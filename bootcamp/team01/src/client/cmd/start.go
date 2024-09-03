package main

import (
	"flag"
	"log"
	"replication/client/client"
	"replication/client/tasks"
	"time"
)

func main() {
	// INITIALIZE
	host := flag.String("H", "", "host of the server")
	port := flag.String("P", "", "port of the server")

	flag.Parse()

	if len(*host) == 0 || len(*port) == 0 {
		flag.PrintDefaults()
		log.Fatalln("host H and port P options are required!")
	}

	// SETUP
	tasksOut := make(chan tasks.Task, 1)
	front := tasks.NewFrontend(tasksOut)
	front.Start()

	_ = client.NewClient(
		client.WithPort(*port),
		client.WithHost(*host),
		client.WithTimeout(time.Second*2),
	)

	// RUN
}
