package models

import "gorm.io/gorm"

const USER_MODEL_NAME = "user"

type User struct {
	gorm.Model
	FullName string  `form:"fullname"`
	Username string  `form:"username"`
	Email    string  `gorm:"index;unique" form:"email"`
	ImageUrl *string `form:"image-url"`
	Password string
	Articles []Article `gorm:"foreignKey:AuthorID"`
	Role     Role      `gorm:"foreignKey:RoleCode;references:Code"`
	RoleCode string
}
