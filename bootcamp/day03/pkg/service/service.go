package service

import (
	"server/pkg/repository"
	"server/types"
	"server/types/models"
)

type Restaurant interface {
	Create(restaurant models.Restaurant) (string, error)
	GetAll(types.PagingIncoming) (types.PaginatedRestaurants, error)
	GetById(id string) (models.Restaurant, error)
	Search(query string, paging types.PagingIncoming) (types.PaginatedRestaurants, error)
	Update(restaurant models.Restaurant) (string, error)
	Delete(id string) (bool, error)
	GetClosest(
		targetId string,
		paging types.PagingIncoming,
		distance int,
	) (types.PaginatedRestaurants, error)
}

type User interface {
	Create(user models.User) (string, error)
	Delete(id string) (bool, error)
	GetById(id string) (models.User, error)
	Update(user models.User) (string, error)
}
type Auth interface {
	LogIn(credentials types.LoginDTO) (string, error)
	LogOut(user models.User) (bool, error)
	Register(user models.User) (string, error)
}

type Service struct {
	Restaurant
	User
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Restaurant: NewRestaurantService(repos.Restaurant),
		User:       NewUserService(repos.User),
		Auth:       NewAuthService(repos.User),
	}
}
