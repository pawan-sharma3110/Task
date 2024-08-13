package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secrectKey = "rest_api"

func GernateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2)})
	return token.SignedString([]byte(secrectKey))
}
