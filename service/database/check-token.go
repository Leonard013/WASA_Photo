package database

import (
	"database/sql"
)

func (db *appdbimpl) CheckToken(token string) error {
	var dummy string
	err := db.c.QueryRow("SELECT userId FROM Users WHERE userId=?", token).Scan(&dummy)
	if err != nil && err != sql.ErrNoRows {
		// An error occurred during the query
		return err
	} else if err == sql.ErrNoRows {
		// No rows were returned
		return err
	}
	return err
}
