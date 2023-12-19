package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Ban(bannerId string, bannedId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT * FROM Ban WHERE bannerId=? AND bannedId=?", bannerId, bannedId).Scan(&dummy)
	if err == nil {
		// A record was found, meaning the user is already banned
		return errors.New("already banned")
	} else if err != nil && err != sql.ErrNoRows {
		// An error occurred other than sql.ErrNoRows
		return err
	}

	_, err = db.c.Exec("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return err
	}
	return nil
}
