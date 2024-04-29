package service

import (
	"candies/server/packages/repository"
	customerrors "candies/server/types/custom_errors"
	"candies/server/types/dtos"
)

type Candy interface {
	BuyCandy(request *dtos.BuyCandyRequestDTO) (dtos.BuyCandySuccessResponseDTO, customerrors.Error)
}

type Service struct {
	Candy
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Candy: NewCandyService(repo.Candy),
	}
}
