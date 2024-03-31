package middleware

import (
	"middleware-prioritas-1/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int, email string) (string, error) {
	//membuat payload
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	//membuat header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//verify signature
	return token.SignedString([]byte(constant.SECRET_JWT))
}