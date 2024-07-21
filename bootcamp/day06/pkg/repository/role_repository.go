package repository

import (
	"github.com/azeek21/blog/models"
	"gorm.io/gorm"
)

type RolesRopo struct {
	db *gorm.DB
}

func NewRolesRepository(db *gorm.DB) RolesRopo {
	return RolesRopo{
		db: db,
	}
}

func (r RolesRopo) GetRoleByRoleCode(role string) (*models.Role, error) {
	res := &models.Role{}
	err := r.db.Take(res, "code = ?", role).Error
	return res, err
}

func (r RolesRopo) GetAllRoles() ([]models.Role, error) {
	roles := []models.Role{}
	err := r.db.Find(roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r RolesRopo) CreateRole(role *models.Role) (*models.Role, error) {
	err := r.db.Create(role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r RolesRopo) DeleteRole(role *models.Role) (bool, error) {
	role.DeletedAt = gorm.DeletedAt{}

	err := r.db.Delete(role).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r RolesRopo) UpdateRole(role *models.Role) (*models.Role, error) {
	err := r.db.Save(role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}
