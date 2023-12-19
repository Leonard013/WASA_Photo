package database

import (
	"database/sql"
	"fmt"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Unfollow(profileId string, followedId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT 1 FROM Follow WHERE profileId=? AND followedId=?", profileId, followedId).Scan(&dummy)
	if err == sql.ErrNoRows {
		// forbidden
		return err

	} else if err != nil && err != sql.ErrNoRows {
		// An error occurred other than sql.ErrNoRows
		fmt.Println("error in unfollow")
		fmt.Println(err)
		return err
	}
	fmt.Println("1")
	_, err = db.c.Exec("DELETE FROM Follow WHERE profileId = ? AND followedId = ?", profileId, followedId)
	return err
}
