package database

func (db *appdbimpl) CheckToken(token string) error {
	err := db.c.QueryRow("SELECT userId FROM Users WHERE userId=?", token).Err()
	return err
}
