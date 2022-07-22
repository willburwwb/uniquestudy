package controller

import (
	"log"
	"test/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("uniquestudy")

func GetToken(name string) (string, error) {
	claim := model.JwtClaim{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	//log.Println("claim", claim)
	//将header和Payload串行转换为字符串
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//log.Println("token", token)
	//假如指定的秘钥，产生签名最终生成令牌
	retoken, err := token.SignedString(secret)
	//log.Println("token sre", retoken, err)
	return retoken, err
}
func ParseToken(tokenString string) (*model.JwtClaim, error) {
	//解码校验，返回*token
	token, err := jwt.ParseWithClaims(tokenString, &(model.JwtClaim{}), func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Println("token解密失败 ", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*model.JwtClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
