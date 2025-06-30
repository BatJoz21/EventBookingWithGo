package models

import (
	"time"

	// "github.com/pelletier/go-toml/query"
	"practice.batjoz/event-booking-with-go/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return nil
}

func (e Event) UpdateEventByID() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var eventSlice []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID)
		if err != nil {
			return nil, err
		}

		eventSlice = append(eventSlice, event)
	}

	return eventSlice, nil
}

func GetEventByID(id int64) (*Event, error) {
	var event Event

	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) DeleteEventByID() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}
