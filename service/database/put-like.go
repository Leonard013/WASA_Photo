package database

import (
	"database/sql"
	"errors"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) PutLike(photoId string, userId string) (string, error) {
	likeId, err := db.CreateLikeId()
	if !errors.Is(err, nil) {
		return "", err
	}
	var dummy string
	err = db.c.QueryRow("SELECT photoId FROM Photos WHERE photoId=?", photoId).Scan(&dummy)
	if !errors.Is(err, nil) {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("photo not found")
		} else {
			return "", err
		}
	}

	err = db.c.QueryRow("SELECT likeId FROM Likes WHERE photoId=? AND author=?", photoId, userId).Scan(&dummy)
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		// An error occurred during the query, return it
		return "", err
	} else if errors.Is(err, nil) {
		// No error occurred, which means a row was found
		return "", errors.New("already liked")
	}

	_, err = db.c.Exec("INSERT INTO Likes (likeId, photoId, author) VALUES (?, ?, ?)", likeId, photoId, userId)
	if !errors.Is(err, nil) {
		return "", err
	}
	return likeId, nil
}
