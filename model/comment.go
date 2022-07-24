package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model        //评论id
	PostID     uint64 //寄给的帖子的id
	Nickname   string //寄者的昵称
	AuthorID   uint64 //寄者的id
	Connent    string
}
