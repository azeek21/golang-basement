package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const USER_CONTEXT_KEY = "user"

type User struct {
	ID        string         `gorm:"primarykey"      json:"id,omitempty"`
	CreatedAt time.Time      `gorm:"<-:create"       json:"createdAt,omitempty"`
	UpdatedAt time.Time      `gorm:"<-:create"       json:"updateAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index;<-:cerate" json:"-"`

	Email    string `gorm:"unique" json:"email"`
	Name     string `              json:"name,omitempty"`
	Username string `gorm:"unique" json:"username,omitempty"`
	Password string `              json:"password"`
}

func (base *User) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.New()
	db.Statement.SetColumn("ID", uuid.String())
	return nil
}
