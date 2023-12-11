package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) GetUser(username string) (string, error) {
	var userId string
	err := db.c.QueryRow("SELECT userId FROM Users WHERE username=?", username).Scan(&userId)
	return userId, err
}
