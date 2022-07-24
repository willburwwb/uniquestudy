package middleware

import (
	"net/http"
	"strings"
	"test/controller/user"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtHeader := context.Request.Header.Get("Authorization")
		if jwtHeader == "" {
			context.JSON(http.StatusBadGateway, gin.H{
				"status":  400,
				"message": "token不存在",
			})
			context.Abort()
			return
		}
		newJwtHeader := strings.Split(jwtHeader, " ")
		if len(newJwtHeader) != 2 || newJwtHeader[0] != "Bearer" {
			context.JSON(http.StatusBadGateway, gin.H{
				"status":  400,
				"message": "token格式错误",
			})
			context.Abort()
			return
		}
		jwtClaims, err := controller.ParseToken(newJwtHeader[1])
		if err != nil || jwtClaims == nil {
			context.JSON(http.StatusBadGateway, gin.H{
				"status":  400,
				"message": "token解析失败或已失效",
			})
			context.Abort()
			return
		}
		context.JSON(http.StatusAccepted, gin.H{
			"status":  200,
			"message": "token解析成功",
		})
	}
}
