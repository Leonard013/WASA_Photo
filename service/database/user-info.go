package database

import "errors"

func (db *appdbimpl) GetUserInfo(userId string) (UserInfo, error) {
	var info UserInfo
	rows, err := db.c.Query("SELECT f.followedId FROM Follow f WHERE f.profileId = ? ", userId)
	if rows.Err() != nil {
		return info, rows.Err()
	}

	if !errors.Is(err, nil) {
		return info, err
	}
	defer rows.Close()
	for rows.Next() {
		var followedId string
		err = rows.Scan(&followedId)
		if !errors.Is(err, nil) {
			return info, err
		}
		info.Following = append(info.Following, followedId)
	}
	rows, err = db.c.Query("SELECT f.profileId FROM Follow f WHERE f.followedId = ? ", userId)
	if rows.Err() != nil {
		return info, rows.Err()
	}

	if !errors.Is(err, nil) {
		return info, err
	}
	defer rows.Close()
	for rows.Next() {
		var profileId string
		err = rows.Scan(&profileId)
		if !errors.Is(err, nil) {
			return info, err
		}
		info.Followers = append(info.Followers, profileId)
	}
	rows, err = db.c.Query("SELECT b.bannedId FROM Ban b WHERE b.bannerId = ? ", userId)
	if rows.Err() != nil {
		return info, rows.Err()
	}

	if !errors.Is(err, nil) {
		return info, err
	}
	defer rows.Close()
	for rows.Next() {
		var bannedId string
		err = rows.Scan(&bannedId)
		if !errors.Is(err, nil) {
			return info, err
		}
		info.Banned = append(info.Banned, bannedId)
	}
	rows, err = db.c.Query("SELECT b.bannerId FROM Ban b WHERE b.bannedId = ? ", userId)
	if rows.Err() != nil {
		return info, rows.Err()
	}

	if !errors.Is(err, nil) {
		return info, err
	}
	defer rows.Close()
	for rows.Next() {
		var bannerId string
		err = rows.Scan(&bannerId)
		if !errors.Is(err, nil) {
			return info, err
		}
		info.IsBanned = append(info.IsBanned, bannerId)
	}
	return info, nil
}
