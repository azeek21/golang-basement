package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	// base
	ID        string         `gorm:"primarykey"      json:"id,omitempty"`
	CreatedAt time.Time      `gorm:"<-:create"       json:"createdAt,omitempty"`
	UpdatedAt time.Time      `gorm:"<-:create"       json:"updateAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index;<-:cerate" json:"-"`

	// data
	Name     string   `json:"name"               csv:"Name"    binding:"required"`
	Address  string   `json:"address"            csv:"Address" binding:"required"`
	Phone    string   `json:"phone"              csv:"Phone"   binding:"required"`
	Location Location `json:"location,omitempty"               binding:"required"`
}

type Location struct {
	// base
	ID        string         `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `                  json:"-"`
	UpdatedAt time.Time      `                  json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"-"`

	// data
	RestaurantId string  `json:"-"`
	Lat          float64 `json:"lat" binding:"required"`
	Lon          float64 `json:"lon" binding:"required"`
}

func (base *Restaurant) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.New()
	db.Statement.SetColumn("ID", uuid.String())
	return nil
}

func (base *Location) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.New()
	db.Statement.SetColumn("ID", uuid.String())
	return nil
}
