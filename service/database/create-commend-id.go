package database

import (
	"github.com/gofrs/uuid"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) CreateCommentId() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	_, err = db.c.Exec("SELECT commendId FROM Comments WHERE commentID=?", id.String())
	if err != nil {
		return db.CreatePhotoId()
	}
	return id.String(), nil
}
