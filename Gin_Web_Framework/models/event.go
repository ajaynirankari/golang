package models

import (
	"fmt"
	"go-lang/Gin_Web_Framework/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)
	`
	fmt.Println("query", query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.DateTime, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	fmt.Println("result", result)
	id, err := result.LastInsertId()
	fmt.Println("result e", e)
	e.Id = id
	fmt.Println("11result e", e)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events = []Event{}

	for rows.Next() {
		var e Event
		err = rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
