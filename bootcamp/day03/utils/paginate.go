package utils

import (
	"gorm.io/gorm"

	"server/types"
)

func Paginate(page types.PagingIncoming) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page.PageNumber - 1) * page.PageSize).Limit(page.PageSize)
	}
}
