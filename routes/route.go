package routes

import (
	"test/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	user := engine.Group("/user")
	{
		user.GET("/login", controller.Login)
		user.POST("/email", controller.GetEmail)
		user.POST("/register", controller.Register)
	}
	return engine
}
