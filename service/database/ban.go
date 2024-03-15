package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Ban(bannerId string, bannedId string) error {
	var dummy string
	var dummy_2 string
	err := db.c.QueryRow("SELECT * FROM Ban WHERE bannerId=? AND bannedId=?", bannerId, bannedId).Scan(&dummy, &dummy_2)
	if errors.Is(err, nil) {
		// A record was found, meaning the user is already banned
		return errors.New("already banned")
	} else if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		// An error occurred other than sql.ErrNoRows
		return err
	}

	_, err = db.c.Exec("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if !errors.Is(err, nil) {
		return err
	}
	return nil
}
