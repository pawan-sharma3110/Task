package models

import (
	"fmt"
	"rest/api/database"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    string `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name,description,location,datetime,user_id) VALUES($1,$2,$3,$4,$5)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return fmt.Errorf("error while scan id %w", err)
	// }
	// e.ID = int(id)
	return nil
}

func GetAllEvent() []Event {
	return events
}
