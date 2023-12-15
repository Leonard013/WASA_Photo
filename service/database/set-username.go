package database

// Update username in the users table
func (db *appdbimpl) SetUsername(username string, userId string) error {
	_, err := db.c.Exec("UPDATE Users SET username = ? WHERE userID = ?", username, userId)
	return err

}
