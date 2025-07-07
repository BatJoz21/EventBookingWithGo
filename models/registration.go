package models

import "practice.batjoz/event-booking-with-go/db"

type Registration struct {
	ID      int64
	EventID int64
	UserID  int64
}

func (r Registration) SaveRegistration() error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	r.ID = id

	return nil
}

func (r Registration) DeleteRegistration() error {
	query := `DELETE FROM registrations WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetRegistrationByID(id int64) (*Registration, error) {
	var regis Registration
	query := `SELECT * FROM registrations WHERE id = ?`
	
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&regis.ID, &regis.EventID, &regis.UserID)
	if err != nil {
		return nil, err
	}

	return &regis, nil
}

func GetAllRegistration() ([]Registration, error) {
	query := `SELECT * FROM registrations`

	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var regisSlice []Registration
	for result.Next() {
		var regis Registration
		err := result.Scan(&regis.ID, &regis.EventID, &regis.UserID)
		if err != nil {
			return nil, err
		}

		regisSlice = append(regisSlice, regis)
	}

	return regisSlice, nil
}
