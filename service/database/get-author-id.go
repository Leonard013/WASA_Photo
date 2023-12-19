package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) GetAuthorId(photoId string) (string, error) {
	var author string
	err := db.c.QueryRow("SELECT author FROM Photos WHERE photoId=?", photoId).Scan(&author)
	return author, err
}
