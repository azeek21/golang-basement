package repository

import (
	"errors"
	"fmt"
	"server/types/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user models.User) (string, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return "0", res.Error
	}
	return user.ID, nil
}

func (r *UserRepository) Delete(id string) (bool, error) {
	res := r.db.Delete(&models.Restaurant{ID: id})
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (r *UserRepository) GetById(id string) (models.User, error) {
	res := models.User{ID: id}
	found := int64(0)
	quey := r.db.Find(&res)
	quey.Count(&found)
	if found == 0 {
		return res, errors.New(fmt.Sprintf("User with id %s not found", id))
	}

	if quey.Error != nil {
		return res, quey.Error
	}
	return res, nil
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	if err := r.db.Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Update(user models.User) (string, error) {
	res := r.db.Save(&user)
	if res.Error != nil {
		return "0", res.Error
	}
	return user.ID, nil
}
