package database

import (
	"errors"

	"github.com/gofrs/uuid"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) CreateUserId() (string, error) {
	id, err := uuid.NewV4()
	if !errors.Is(err, nil) {
		return "", err
	}
	_, err = db.c.Exec("SELECT * FROM Users WHERE userId=?", id.String())
	if !errors.Is(err, nil) {
		return db.CreateUserId()
	}
	return id.String(), nil
}
