package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Vote     uint64
	AuthorID uint64
	Title    string `form:"title" binding:"required"`
	Content  string `form:"content" binding:"required"`
	Nickname string
}
