package models

import "gorm.io/gorm"

const ROLE_MODEL_NAME = "role"

type Role struct {
	gorm.Model
	Code        string `gorm:"unique;index" json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
