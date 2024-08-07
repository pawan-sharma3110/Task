package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbIn() *sql.DB {
	conStr := `host=localhost port=5432 user=postgres dbname=task password=Pawan@2003 sslmode=disable`
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
