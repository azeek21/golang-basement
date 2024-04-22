package repository

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"

	"server/types"
	"server/types/models"
)

type Restaurant interface {
	Create(restaurant models.Restaurant) (string, error)
	GetAll(pagination types.PagingIncoming) (types.PaginatedRestaurants, error)
	GetById(id string) (models.Restaurant, error)
	Search(query string, paging types.PagingIncoming) (types.PaginatedRestaurants, error)
	Update(restaurant models.Restaurant) (string, error)
	Delete(id string) (bool, error)
	GetClosest(
		target models.Restaurant,
		paging types.PagingIncoming,
		distance int,
	) (types.PaginatedRestaurants, error)
}

type User interface {
	Create(user models.User) (string, error)
	Delete(id string) (bool, error)
	GetById(id string) (models.User, error)
	Update(user models.User) (string, error)
	GetByEmail(email string) (models.User, error)
}

type Repository struct {
	Restaurant
	User
}

func NewRepository(db *gorm.DB, elastic *elasticsearch.Client) *Repository {
	return &Repository{
		Restaurant: NewRestaurantsRepository(db, elastic),
		User:       NewUserRepository(db),
	}
}
