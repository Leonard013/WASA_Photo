package database

func (db *appdbimpl) IfBanned(bannerId string, bannedId string) error {
	err := db.c.QueryRow("SELECT * FROM Ban WHERE bannerId=?, bannedId=?", bannerId, bannedId).Err()
	if err != nil {
		return nil
	}
	return err
}
