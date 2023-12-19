package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) Uncomment(commentId string, userId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT 1 FROM Comments WHERE commentId=? AND author=?", commentId, userId).Scan(&dummy)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM Comments WHERE commentId=? AND author=?", commentId, userId)
	return err
}
