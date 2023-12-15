package database

import "errors"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Unban(bannerId string, bannedId string) error {
	err := db.c.QueryRow("SELECT * FROM Ban WHERE bannerId=?, bannedId=?", bannerId, bannedId).Err()
	if err == nil {
		// forbidden
		return errors.New("already unbanned")

	}

	_, err = db.c.Exec("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	return err
}
