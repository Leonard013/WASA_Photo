package database

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) CreatePhotoId() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	var dummy string
	err = db.c.QueryRow("SELECT photoId FROM Photos WHERE photoId=?", id.String()).Scan(&dummy)
	if err == sql.ErrNoRows {
		// The ID does not exist in the database, so it's safe to use
		return id.String(), nil
	} else if err != nil && err != sql.ErrNoRows {
		// An error occurred that is not sql.ErrNoRows
		return "", err
	}
	return db.CreatePhotoId()

}
