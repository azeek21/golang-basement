package service

import (
	"server/pkg/repository"
	"server/types/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(user models.User) (string, error) {
	return s.repo.Create(user)
}

func (s *UserService) Delete(id string) (bool, error) {
	return s.repo.Delete(id)
}

func (s *UserService) GetById(id string) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Update(user models.User) (string, error) {
	return s.repo.Update(user)
}
