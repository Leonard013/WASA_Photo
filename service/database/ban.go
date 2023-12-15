package database

import "errors"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Ban(bannerId string, bannedId string) error {
	err := db.c.QueryRow("SELECT * FROM Ban WHERE bannerId=?, bannedId=?", bannerId, bannedId).Err()
	if err == nil {
		// forbidden
		return errors.New("already banned")

	}
	_, err = db.c.Exec("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return err
	}
	return nil
}
