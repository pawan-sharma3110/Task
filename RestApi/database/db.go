package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DbIn() {
	var err error
	connStr := `host=localhost port=5432 user=postgres dbname=task password=Pawan@2003 sslmode=disable`
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Could not ping database:", err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	// Uncomment this if you need to create tables
	createTable(DB)
}

func createTable(db *sql.DB) {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	  id SERIAL PRIMARY KEY,
	  email TEXT NOT NULL UNIQUE,
	  password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    datetime TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
)`

	_, err = db.Exec(createEventsTable)
	if err != nil {
		// panic("Could not create events table")
		log.Fatal(err)
	}
}
