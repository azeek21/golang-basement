package service

import (
	"server/pkg/repository"
	"server/types"
	"server/types/models"
)

type RestaurantService struct {
	repo repository.Restaurant
}

func NewRestaurantService(repo repository.Restaurant) *RestaurantService {
	return &RestaurantService{
		repo: repo,
	}
}

func (s *RestaurantService) Create(restaurant models.Restaurant) (string, error) {
	return s.repo.Create(restaurant)
}

func (s *RestaurantService) Delete(id string) (bool, error) {
	return s.repo.Delete(id)
}

func (s *RestaurantService) GetAll(
	pagination types.PagingIncoming,
) (types.PaginatedRestaurants, error) {
	return s.repo.GetAll(pagination)
}

func (s *RestaurantService) GetById(id string) (models.Restaurant, error) {
	return s.repo.GetById(id)
}

func (s *RestaurantService) Search(
	query string,
	paging types.PagingIncoming,
) (types.PaginatedRestaurants, error) {
	return s.repo.Search(query, paging)
}

func (s *RestaurantService) Update(
	restaurant models.Restaurant,
) (string, error) {
	return s.repo.Update(restaurant)
}

func (s *RestaurantService) GetClosest(
	targetId string,
	paging types.PagingIncoming,
	distance int,
) (types.PaginatedRestaurants, error) {
	res := types.PaginatedRestaurants{}
	target, err := s.GetById(targetId)
	if err != nil {
		return res, err
	}
	return s.repo.GetClosest(target, paging, distance)
}
