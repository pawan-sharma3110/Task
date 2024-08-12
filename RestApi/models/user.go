package models

import (
	"errors"
	"rest-api/database"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() (*int64, error) {
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	var id int64
	err = database.DB.QueryRow(query, u.Email, hash).Scan(&id)
	if err != nil {
		return nil, err
	}
	u.ID = id
	return &u.ID, nil
}
func (u User) ValidateCredentials() error {
	var retrievedPass string
	query := `SELECT password FROM users WHERE email=$1`
	err := database.DB.QueryRow(query, u.Email).Scan(&retrievedPass)
	if err != nil {
		return errors.New("credentials invalid email")
	}
	passwordValid := utils.CompairePAssword(u.Password, retrievedPass)
	if !passwordValid {
		return errors.New("credentials invalid")
	}
	return nil
}
