package service

import (
	"candies/server/packages/repository"
	customerrors "candies/server/types/custom_errors"
	"candies/server/types/dtos"
	"fmt"
	"net/http"
)

type CandyService struct {
	repo repository.Candy
}

func NewCandyService(repo repository.Candy) *CandyService {
	return &CandyService{
		repo: repo,
	}
}

func (candservice *CandyService) BuyCandy(request *dtos.BuyCandyRequestDTO) (dtos.BuyCandySuccessResponseDTO, customerrors.Error) {
	respp := dtos.BuyCandySuccessResponseDTO{Thanks: "Thanks for your purchase."}
	candy, err := candservice.repo.GetByName(request.CandyType)

	if err != nil {
		return respp, customerrors.NewCustomError(http.StatusBadRequest, err.Error())
	}

	total := candy.Price * request.CandyCount

	if total > request.Money {
		return respp, customerrors.NewCustomError(http.StatusPaymentRequired, fmt.Sprintf("no enought money to buy %d %s", request.CandyCount, candy.Name))
	}

	respp.Change = request.Money - total
	return respp, nil
}
