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

func GetAllEvent() ([]Event, error) {
	query := `SELECT id, name, location, datetime, user_id FROM events`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing statement: %w", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return events, nil
}
