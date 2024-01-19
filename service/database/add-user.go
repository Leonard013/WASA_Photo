package database

import "errors"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) AddUser(username string) (string, error) {
	id, err := db.CreateUserId()
	if !errors.Is(err, nil) {
		return "", err
	}
	_, err = db.c.Exec("INSERT INTO Users (userId, username) VALUES (?, ?)", id, username)
	return id, err
}
