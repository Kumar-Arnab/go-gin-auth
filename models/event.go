package models

import (
	"time"

	"github.com/Kumar-Arnab/events-rests-auth/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() (Event, error) {
	// ? is a sql injection safe way of injecting values in this query
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return Event{}, err
	}
	defer stmt.Close()

	// Exec() is also a Query() method
	// If we have a method that changes stuff, its Exec()
	// If we have a method that queries db, its Query()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return Event{}, err
	}
	id, err := result.LastInsertId()

	e.ID = id
	return e, err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// Here also 1st we can prepare a Prepare() method but since its a normal query we can directly use Query()
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event

		// populating event struct with the data in the row by passing a reference of event fields to row.Scan()
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := `
		DELETE FROM events
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	return err
}
