package models

import (
	"errors"

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

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	// populating retrivedPassword with the password returned by the DB
	var retrivedPassword string
	err := row.Scan(&user.ID, &retrivedPassword)

	if err != nil {
		return errors.New("Invalid credentials 1")
	}

	valid := utils.CheckPasswordHash(user.Password, retrivedPassword)

	if !valid {
		return errors.New("Invalid credentials 2")
	}

	return nil
}
