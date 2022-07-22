package controller

import (
	"log"
	"net/http"
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
)

func GetLostPasswd(context *gin.Context) {
	rdb := database.GetRdb()
	db := database.GetDB()
	email := context.PostForm("email")
	code := context.PostForm("code")
	ncode, err := rdb.Get(email).Result()

	if err != nil {
		log.Println("redis 错误", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "redis 错误",
		})
		return
	}
	if ncode == "" || code != ncode {
		log.Println("验证码失效或错误")
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "验证码错误",
		})
		return
	}
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		log.Println("还未注册")
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "未注册",
		})
		return
	}
	password, err := Decode(user.Password)
	if err != nil {
		log.Println("解密错误", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "解密错误",
		})
		return
	}
	log.Println("解密成功")
	context.JSON(http.StatusOK, gin.H{
		"status": 200,
		"passwd": string(password),
	})
}
