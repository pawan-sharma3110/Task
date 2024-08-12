package models

import (
	"fmt"
	"rest-api/database"
	"time"
)

type Event struct {
	ID          int64
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
	e.ID = int64(id)
	e.DateTime = time.Now()
	fmt.Print(id)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name,description, location, datetime, user_id FROM events`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing statement: %w", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
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
func GetEventById(id int64) (*Event, error) {
	var event Event
	query := `SELECT * FROM events WHERE id =$1`
	err := database.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name =$2,description=$3,location=$4,dateTime=$5,user_id=$6
		WHERE id =$1`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.Location, time.Now(), event.UserID)
	return err
}
func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id =$1`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}
