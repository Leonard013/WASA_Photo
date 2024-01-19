package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Unban(bannerId string, bannedId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT 1 FROM Ban WHERE bannerId=? AND bannedId=?", bannerId, bannedId).Scan(&dummy)
	if errors.Is(err, sql.ErrNoRows) {
		// forbidden
		return err
	} else if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		// An error occurred other than sql.ErrNoRows
		return err
	}

	_, err = db.c.Exec("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	return err
}
