package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	// bcrypt package from crypto package has this function to generate hashed password
	// it accepts 2 parameters: byte[] of string and weight which simply controls how complex the hashing will be
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}
