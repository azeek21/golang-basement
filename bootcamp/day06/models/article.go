package models

import (
	"gorm.io/gorm"
)

const ARTICLE_MODEL_NAME = "article"

type Article struct {
	gorm.Model
	Title       string  `form:"title" binding:"required"`
	Description string  `form:"description" binding:"required"`
	Content     string  `form:"content" binding:"required"`
	AuthorID    uint    `form:"authorId"`
	ImageUrl    *string `form:"image"`
}

func (a Article) HasImage() bool {
	return a.ImageUrl != nil && len(*a.ImageUrl) > 0
}

func (a Article) GetImage() string {
	if a.HasImage() {
		return *a.ImageUrl
	}
	return ""
}
