package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) DeletePhoto(photoId string, userId string) (string, error) {
	var author, path string
	err := db.c.QueryRow("SELECT author,photoPath FROM Photos WHERE photoId=?", photoId).Scan(&author, &path)
	if err != nil && err != sql.ErrNoRows {
		// An error occurred during the query, return it
		return "", err
	} else if err == sql.ErrNoRows {
		// No error occurred, which means a row was found
		return "", errors.New("already deleted")
	}

	if author != userId {
		// forbidden
		return "", errors.New("not your photo")
	}

	_, err = db.c.Exec("DELETE FROM Photos WHERE photoId=? AND author=?", photoId, userId)
	return path, err
}
