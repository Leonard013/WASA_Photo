package database

import "errors"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Follow(profileId string, followedId string) error {
	err := db.c.QueryRow("SELECT * FROM Follow WHERE profileId=?, followedId=?", profileId, followedId).Err()
	if err == nil {
		// forbidden
		return errors.New("already following")

	}
	_, err = db.c.Exec("INSERT INTO Follow (profileId, followedId) VALUES (?, ?)", profileId, followedId)
	if err != nil {
		return err
	}
	return nil
}
