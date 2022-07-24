package routes

import (
	ccomment "test/controller/comment"
	cpost "test/controller/post"
	cuser "test/controller/user"

	"test/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	user := engine.Group("/user")
	{
		user.POST("/login", cuser.Login)
		user.POST("/email", cuser.GetEmail)
		user.POST("/register", cuser.Register)
		user.POST("/getLostPasswd", cuser.GetLostPasswd)
	}
	post := engine.Group("/post")
	{
		post.POST("/create", middleware.JWT(), cpost.CreatPost)
		post.DELETE("/delete/:id", middleware.JWT(), cpost.DeletePost)
		post.GET("/getPostsByTitle", middleware.JWT(), cpost.GetPostsByTitle)
		post.GET("/getPostsByTime", middleware.JWT(), cpost.GetPostsByTime)
		post.PUT("/updatePost", middleware.JWT(), middleware.Judge(), cpost.UpdatePost)
		post.GET("/getPostByVote", middleware.JWT(), cpost.GetPostByVote)
	}
	comment := engine.Group("/comment")
	{
		comment.POST("/create", middleware.JWT(), ccomment.Create)
		comment.GET("/getCommentList", middleware.JWT(),ccomment.GetCommentList)
		comment.DELETE("/delete", middleware.JWT(),ccomment.Delete)
	}
	return engine
}
