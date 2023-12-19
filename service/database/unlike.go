package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Unlike(photoId string, userId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT 1 FROM Likes WHERE photoId=? AND author=?", photoId, userId).Scan(&dummy)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM Likes WHERE photoId = ? AND author=?", photoId, userId)
	return err
}
