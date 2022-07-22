package routes

import (
	"log"
	"test/controller"
	"test/middleware"

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

	engine.POST("/", middleware.JWT(), func(ctx *gin.Context) {
		log.Println("中间件检验   成功")
	})
	return engine
}
