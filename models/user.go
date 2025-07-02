package models

import (
	"errors"

	"practice.batjoz/event-booking-with-go/db"
	"practice.batjoz/event-booking-with-go/utils"
)

type Users struct {
	User_ID  int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) SaveUser() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	encryptedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, encryptedPass)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.User_ID = id

	return nil
}

func (u Users) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&retrivedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
