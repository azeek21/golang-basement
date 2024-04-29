package repository

import (
	"candies/server/types/models"
	"errors"
)

type CandyRepository struct {
}

func NewCandyRepository() *CandyRepository {
	return &CandyRepository{}
}

func (candyrepository *CandyRepository) Create() (string, error) {
	return "", nil
}

func (candyrepository *CandyRepository) GetById() error {
	return nil
}
func (candyrepository *CandyRepository) GetByName(name string) (models.Candy, error) {
	candy := models.CANDIES[name]
	if candy == nil {
		return models.Candy{}, errors.New("No such candy with name " + name)
	}
	return *candy, nil
}
func (candyrepository *CandyRepository) Delete() (bool, error) {
	return false, nil
}
