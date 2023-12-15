package database

import "errors"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Unfollow(profileId string, followedId string) error {
	err := db.c.QueryRow("SELECT * FROM Follow WHERE profileId=?, followedId=?", profileId, followedId).Err()
	if err == nil {
		// forbidden
		return errors.New("already unfollowed")

	}

	_, err = db.c.Exec("DELETE FROM Follow WHERE profileId = ? AND followedId = ?", profileId, followedId)
	return err
}
