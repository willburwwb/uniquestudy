package routes

import (
	cpost "test/controller/post"
	cuser "test/controller/user"
	"test/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	user := engine.Group("/user")
	{
		user.GET("/login", cuser.Login)
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
		post.PUT("/updatePost", middleware.JWT(), cpost.UpdatePost)
		post.GET("/getPostByVote", middleware.JWT(), cpost.GetPostByVote)
	}
	return engine
}
