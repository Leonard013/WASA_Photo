package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Comment(photoId string, userId string, text string, t string) (string, error) {
	commentId, err := db.CreateLikeId()
	if err != nil {
		return "", err
	}

	var dummy string
	err = db.c.QueryRow("SELECT 1 FROM Photos WHERE photoId=?", photoId).Scan(&dummy)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("photo does not exist")
		} else {
			return "", err
		}
	}

	_, err = db.c.Exec("INSERT INTO Comments (commentId, photoId, author, text, date) VALUES (?, ?,?,?,?)", commentId, photoId, userId, text, t)
	return commentId, err
}
