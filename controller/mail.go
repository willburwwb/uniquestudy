package controller

import (
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"test/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randCode() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func SendEmail(userEmail string) (string, error) {
	e := email.NewEmail()
	e.From = "wwb <709782717@qq.com>"
	e.To = []string{userEmail}
	e.Subject = "uniquestudy gin框架邮箱测试"
	randcode := randCode()
	e.Text = []byte("验证码为" + randcode)
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "709782717@qq.com", "xyhxfdulxgwzbbed", "smtp.qq.com"))
	if err != nil {
		log.Println("email send to ", userEmail, "failed && randcode =", randcode)
	}
	return randcode, err
}
func GetEmail(context *gin.Context) {
	email := context.PostForm("email")
	randcode, err := SendEmail(email)
	//是否可发送验证码，并将randCode放入redis中保存
	if err != nil {
		log.Println("Send email failed ", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	rdb := database.GetRdb()
	err = rdb.Set(email, randcode, time.Minute*10).Err()
	if err != nil {
		log.Println("Set redis failed", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	log.Println("Send email successfully")
	context.JSON(http.StatusAccepted, gin.H{
		"status":   200,
		"randcode": randcode,
	})
}
