package middleware

import (
	"net/http"
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
)

func Tudge() gin.HandlerFunc {
	return func(context *gin.Context) {
		postid := context.Param("id")
		username := context.PostForm("name")
		db := database.GetDB()
		var post model.Post
		var user model.User
		db.Where("id=?", postid).First(&post)
		db.Where("id=?", post.AuthorID).First(&user)
		if user.Name != username {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "该用户无权限",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
