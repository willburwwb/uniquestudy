package middleware

import (
	"net/http"
	"strconv"
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
)

func Judge() gin.HandlerFunc {
	return func(context *gin.Context) {
		postid, _ := strconv.Atoi(context.PostForm("id"))
		username := context.PostForm("name")
		db := database.GetDB()
		var post model.Post
		var user model.User
		db.Where("id=?", postid).First(&post)
		db.Where("id=?", post.AuthorID).First(&user)
		//log.Println(postid, " user.", user.Name, " username", username)
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
