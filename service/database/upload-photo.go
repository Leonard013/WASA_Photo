package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) UploadPhoto(photoId string, title string, photoPath string, userId string, time string) error {
	_, err := db.c.Exec("INSERT INTO Photos (photoId, title, photoPath, date, author) VALUES (?, ?,?,?,?)", photoId, title, photoPath, time, userId)
	return err
}
