package utils

import (
	"errors"
	"fmt"
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
func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexcepted sining methord")
		}
		return secrectKey, nil
	})
	if err != nil {
		return errors.New("could not parsed token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}
	email := claims["email"].(string)
	userId := claims["userId"].(int64)
	fmt.Println(email, userId)
	return nil
}
