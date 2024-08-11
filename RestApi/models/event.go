package models

import (
	"log"
	"rest/api/database"
	"time"
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

func (e Event) Save() error {
	query := `INSERT INTO events(name,description,location,datetime,user_id )VALUES($1,$2,$3,$4,$5)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, time.Now(), e.UserID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// id, err := result.LastInsertId()
	// e.ID = int(id)
	return err
}

func GetAllEvent() []Event {
	return events
}
