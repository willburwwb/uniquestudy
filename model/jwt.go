package model

import "github.com/dgrijalva/jwt-go"

type JwtClaim struct {
	Name string
	jwt.StandardClaims
}
