package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"randomaliens/internal/handler"
	"randomaliens/internal/repository"
	"randomaliens/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context, port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.Printf("Tcp listener created, litening at: %s", listener.Addr())

	if err != nil {
		return err
	}

	repo := repository.NewRepository(nil)
	services := service.NewService(repo)

	grpcServer := grpc.NewServer()
	handler := handler.NewHandler(services)
	handler.Init(grpcServer)

	log.Println("Transmitter server registered")

	go func() {
		log.Printf("Starting grpc server at %s\tServerInfo: %+v\n", listener.Addr(), grpcServer.GetServiceInfo())
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

	log.Println("Registered interrupt listener for grpc server")
	select {
	case <-ctx.Done():
		log.Print("Grpc server shutting down")
		grpcServer.GracefulStop()
		return nil
	}
	return nil
}
