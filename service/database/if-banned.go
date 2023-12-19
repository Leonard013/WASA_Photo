package database

func (db *appdbimpl) IfBanned(bannerId string, bannedId string) error {
	var dummy string
	err := db.c.QueryRow("SELECT * FROM Ban WHERE bannerId=? AND bannedId=?", bannerId, bannedId).Scan(&dummy)
	return err
}
