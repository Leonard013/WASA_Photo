package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Follow(profileId string, followedId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT * FROM Follow WHERE profileId=? AND followedId=?", profileId, followedId).Scan(&dummy)
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		// An error occurred during the query, return it
		return err
	} else if errors.Is(err, nil) {
		// No error occurred, which means a row was found
		return errors.New("already following")
	}
	// if err != sql.ErrNoRows
	_, err = db.c.Exec("INSERT INTO Follow (profileId, followedId) VALUES (?, ?)", profileId, followedId)
	if !errors.Is(err, nil) {
		return err
	}
	return nil
}
