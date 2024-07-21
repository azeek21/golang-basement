package service

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{
		repo: userRepo,
	}
}

func (s userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s userService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.CreateUser(user)
}
func (s userService) GetUserById(id uint) (*models.User, error) {
	return s.repo.GetUserById(id)
}
func (s userService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}
func (s userService) DeleteUser(user *models.User) (bool, error) {
	return s.repo.DeleteUser(user)
}
func (s userService) SetRole(user *models.User, role string) (*models.User, error) {
	return s.repo.SetRole(user, role)
}
func (s userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}
