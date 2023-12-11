package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) AddUser(id string, username string) error {
	_, err := db.c.Exec("INSERT INTO Users (userId, username) VALUES (?, ?)", id, username)
	return err
}
