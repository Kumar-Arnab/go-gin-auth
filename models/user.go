package models

import (
	"github.com/Kumar-Arnab/events-rests-auth/db"
	"github.com/Kumar-Arnab/events-rests-auth/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user User) Save() (User, error) {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	passwordHash, err := utils.HashPassword(user.Password)

	if err != nil {
		return User{}, err
	}

	result, err := stmt.Exec(user.Email, passwordHash)

	if err != nil {
		return User{}, err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	user.Password = passwordHash
	return user, err
}
