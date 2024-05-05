package handler

import (
	"randomaliens/internal/grpc/transmitter"
	"randomaliens/internal/service"

	"google.golang.org/grpc"
)

type Handler struct {
	service *service.Service
	transmitter.UnimplementedTransmitterServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) Init(srv *grpc.Server) error {
	transmitter.RegisterTransmitterServer(srv, h)
	return nil
}
