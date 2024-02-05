package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "some-secret-key"

// function to generate token
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

func ValidateToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// checking the signing algo used to encrypt the token
		// .(____) is a method for checking in go for a specific type
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}

	/*
		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		if !ok {
			return errors.New("Invalid token claims.")
		}

		// jwt.MapClaims values are essentially a map type
		// the fields name are same as the field name we used to set claims in our previous function
		// also we can match the types using go's .() specific type check property
		email := claims["email"].(string)
		userId := claims["userId"].(int64)
	*/

	return nil
}
