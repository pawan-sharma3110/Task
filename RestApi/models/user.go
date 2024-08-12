package models

import (
	"rest/api/database"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() (*int64, error) {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	var id int64
	err := database.DB.QueryRow(query, u.Email, u.Password).Scan(&id)
	if err != nil {
		return nil, err
	}
	u.ID = id
	return &u.ID, nil
}
