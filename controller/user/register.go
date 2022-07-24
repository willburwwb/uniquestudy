package controller

import (
	"log"
	"net/http"
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func Register(context *gin.Context) {
	db := database.GetDB()
	rdb := database.GetRdb()
	name := context.PostForm("name")
	password := context.PostForm("password")
	email := context.PostForm("email")
	randcode := context.PostForm("randcode")

	//对姓名，邮箱，密码进行简单的校验

	//检测验证码是否合法
	if !checkCode(rdb, email, randcode) {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "验证码不存在. Please test again",
		})
		return
	}

	//检验是否已经存在该邮箱
	if isEmailExist(db, email) {
		context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "the email has been used",
		})
		return
	}

	base64Password := Encode(password)

	user := model.User{
		Name:     name,
		Password: base64Password,
		Email:    email,
	}

	db.Create(&user)
	log.Println("insert a new user", name, base64Password, email)
	context.JSON(http.StatusAccepted, gin.H{
		"status":  200,
		"message": "insert a new user",
	})

}
func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email=?", email).First(&user)
	return user.ID != 0
}
func checkCode(rdb *redis.Client, email string, randcode string) bool {
	val, err := rdb.Get(email).Result()
	if err != nil {
		log.Println("获取redis错误 ", err)
		return false
	}
	if val == "" {
		log.Println("验证码不存在 ", err)
		return false
	}
	if val != randcode {
		log.Println("验证码错误 ", err)
		return false
	}
	return true
}
