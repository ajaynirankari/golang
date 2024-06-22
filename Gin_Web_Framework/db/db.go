package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Println("Error connecting to database:", err)
		return err
	}
	err = DB.Ping()
	if err != nil {
		log.Println("Error pinging database:", err)
		return err
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	return createTable()
}

func createTable() error {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER
)
	`
	_, err := DB.Exec(createEventTable)

	if err != nil {
		log.Println("Error creating table event:", err)
		return err
	}
	return nil
}
