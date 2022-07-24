package comment

import (
	"log"
	"net/http"
	"strconv"
	controller "test/controller/post"
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context) {
	var comment model.Comment
	var nick model.Nick
	var nickname string
	if err := context.ShouldBind(&comment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "参数错误",
		})
		return
	}
	authorid, err := strconv.Atoi(context.PostForm("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "id格式错误",
		})
		return
	}
	db := database.GetDB()
	db.Where("author_id=? AND post_id=?", authorid, comment.PostID).Find(&nick)
	if nick.Nickname == "" {
		nickname = controller.SetNickname()
		nick.Nickname = nickname
		nick.AuthorID = uint64(authorid)
		nick.PostID = comment.PostID
		db.Create(nick)
		log.Println("创建新的匿名对象", nickname)
	} else {
		nickname = nick.Nickname
	}
	comment.Nickname = nickname
	db.Create(&comment)
	log.Println("创建评论成功")
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "评论创建成功",
	})
}
func Delete(context *gin.Context) {
	var commentDelete model.CommentDelete
	if err := context.ShouldBind(&commentDelete); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "参数错误",
		})
		return
	}
	db := database.GetDB()
	db.Where("id=?", commentDelete.CommentID).Delete(&model.Comment{})
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "评论创建成功",
	})
}
func GetCommentList(context *gin.Context) {
	var commentList model.CommentList
	if err := context.ShouldBind(&commentList); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": "参数错误",
		})
	}
	db := database.GetDB()
	var comments []model.Comment
	db.Where("post_id=?", commentList.PostID).Find(&comments)
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": comments,
	})
}
