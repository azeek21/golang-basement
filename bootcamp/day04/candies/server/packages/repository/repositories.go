package repository

import "candies/server/types/models"

type Candy interface {
	Create() (string, error)
	GetById() error
	GetByName(name string) (models.Candy, error)
	Delete() (bool, error)
}

type Repository struct {
	Candy
}

func NewRepository() *Repository {
	return &Repository{
		Candy: NewCandyRepository(),
	}
}
