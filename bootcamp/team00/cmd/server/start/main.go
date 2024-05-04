package main

import (
	"fmt"
	"log"
	"net"
	"randomaliens/server/pkgs/handlers"
	grpc "google.golang.org/grpc"
	reandomaliens "randomalies/server/pkgs/services/transmitter"
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf(err.Error())
	}

	handler := handlers.Handlers{}
	grpcServer := grpc.NewServer()
	transmitter.
}
