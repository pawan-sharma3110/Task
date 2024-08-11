package models

import (
	"fmt"
	"rest/api/database"
	"time"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name,description,location,datetime,user_id) VALUES($1,$2,$3,$4,$5)RETURNING id`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()
	var id int
	err = stmt.QueryRow(e.Name, e.Description, e.Location, time.Now(), e.UserID).Scan(&id)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
	e.ID = id
	e.DateTime = time.Now()
	fmt.Print(id)
	return nil
}

func GetAllEvent() []Event {
	return events
}
