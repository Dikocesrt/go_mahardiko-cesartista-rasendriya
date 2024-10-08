package middleware

import (
	"clean/constants"
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
	return token.SignedString([]byte(constants.SECRET_JWT))
}