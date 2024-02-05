package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "some-secret-key"

func GenerateToken(email string, userId int64) (string, error) {

	// this function generates a token which simple means with data that's attached to it
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour).Unix(),
		})

	return token.SignedString([]byte(secretKey))
}
