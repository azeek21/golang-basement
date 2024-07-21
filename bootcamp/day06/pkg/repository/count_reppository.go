package repository

import "gorm.io/gorm"

type CountRepo struct {
	db *gorm.DB
}

func NewCountRepository(db *gorm.DB) CountRepository {
	return CountRepo{
		db: db,
	}
}

func (r CountRepo) Count(model interface{}) int64 {
	res := int64(0)
	err := r.db.Model(model).Count(&res).Error
	if err != nil {
		return 0
	}
	return res
}
