package service

import (
	"randomaliens/internal/grpc/transmitter"
	"randomaliens/internal/repository"
)

type Generator interface {
	Uuid() string
	// Mean and Deviation must be randomly selected within the two given arguments
	Mean(float64, float64) float64
	Deviation(float64, float64) float64
	Frequency(mean float64, deviation float64) float64 // frequency is generated depending on Deviation and Mean
}

type Session interface {
	SaveAnomaly(record *transmitter.Record, sesison *AnomalySession) error
	NewSession(id string, threshold float64) (*AnomalySession, error)
	SaveSession(sesison *AnomalySession) error
}

type Service struct {
	Generator
	Session
}

func NewService(repo *repository.Repository) *Service {

	return &Service{
		Generator: NewGeneratorService(),
		Session:   NewSessionSevice(repo),
	}
}
