package controller

import (
	"log"
	"net/http"
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	db := database.GetDB()
	name := context.PostForm("name")
	password := context.PostForm("password")
	password = Encode(password)
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		log.Println("user isn't exsit")
		context.JSON(http.StatusBadGateway, gin.H{
			"status":  400,
			"message": "用户不存在",
		})
		return
	}
	if user.Password == password {
		log.Println("login successfully")
		context.JSON(http.StatusAccepted, gin.H{
			"status":  200,
			"message": "login successfully",
		})
		token, err := GetToken(user.Name)
		if err != nil {
			log.Println("token 为", token)
			log.Println("token 未成功生成", err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}
	log.Println("密码错误")
	context.JSON(http.StatusBadGateway, gin.H{
		"status":  200,
		"message": "密码错误",
	})
}
