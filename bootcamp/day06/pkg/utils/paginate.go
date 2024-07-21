package utils

import (
	"github.com/azeek21/blog/models"
	"gorm.io/gorm"
)

func Paginate(paging models.PagingIncoming) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (paging.Page - 1) * paging.ItemsPerPage
		return db.Offset(offset).Limit(paging.ItemsPerPage)
	}
}
