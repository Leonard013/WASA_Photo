package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUsername(id string) (string, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM Users WHERE userId=?", id).Scan(&username)
	return username, err
}
