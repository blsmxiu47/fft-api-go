package main

import (
	"database/sql"
	"errors"
)

type user struct {
	ID int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstname"`
}

func (u* user) getUser(db* sql.DB) error {
	return errors.New("this is an expected error #tktk")
}
