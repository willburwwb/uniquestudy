package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model        //评论id
	PostID     uint64 `form:"postID" binding:"required"` //寄给的帖子的id
	ResponseID uint64 `form:"ResponseID" `               //回复的评论ID
	Nickname   string //寄者的昵称
	Connent    string `form:"connent" binding:"required"` //评论的内容
}

//不同的帖子下的人的昵称
type Nick struct {
	Nickname string
	AuthorID uint64
	PostID   uint64
}

type CommentDelete struct {
	ID        uint64 `form:"id" binding:"required"`        //删除者的id
	CommentID uint64 `form:"commentID" binding:"required"` //要删除的id
}

type CommentList struct {
	ID     uint64 `uri:"id" form:"id" binding:"required"`
	PostID uint64 `uri:"postID" form:"postID" binding:"required"`
}
