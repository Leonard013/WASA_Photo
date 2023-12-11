package database

import (
	"github.com/gofrs/uuid"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) CreateId(category string) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	_, err = db.c.Exec("SELECT userId FROM ? WHERE userID=?", category, id.String())
	if err == nil {
		return db.CreateId(category)
	}
	return id.String(), nil
}
