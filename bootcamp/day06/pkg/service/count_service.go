package service

import (
	"github.com/azeek21/blog/pkg/repository"
)

type CountServ struct {
	repo repository.CountRepository
}

func NewCountService(repo repository.CountRepository) CountingService {
	return CountServ{
		repo: repo,
	}
}

func (s CountServ) Count(model interface{}) int {
	return int(s.repo.Count(model))
}
