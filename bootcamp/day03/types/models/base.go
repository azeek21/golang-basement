package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `                  json:"createdAt"`
	UpdatedAt time.Time      `                  json:"updateAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"-"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.New()
	db.Statement.SetColumn("ID", uuid.String())
	return nil
}
