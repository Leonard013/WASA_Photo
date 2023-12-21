package database

import (
	"database/sql"
	"os"
	"sort"
	"time"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) GetStream(userId string) ([][]byte, error) {
	var following []string
	var images_stream []Photo
	row, err := db.c.Query("SELECT followedId FROM Follow WHERE profileId = ?", userId)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var followedId string
		err = row.Scan(&followedId)
		if err != nil {
			return nil, err
		}

		var dummy string
		err := db.c.QueryRow("SELECT 1 FROM Ban WHERE bannerId = ? AND bannedId = ?", followedId, userId).Scan(&dummy)
		if err != nil {
			if err == sql.ErrNoRows {
				following = append(following, followedId)
			} else {
				return nil, err
			}
		}
	}

	for _, followedId := range following {
		var image []byte
		var path string
		var date string
		err := db.c.QueryRow("SELECT photoPath,date FROM Photos WHERE author = ?", followedId).Scan(&path, &date)
		if err != nil {
			return nil, err
		}
		img, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer img.Close()
		t, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			return nil, err
		}
		photo := Photo{
			Image: image,
			Date:  t,
		}
		images_stream = append(images_stream, photo)
	}

	sort.Slice(images_stream, func(i, j int) bool {
		return images_stream[i].Date.Before(images_stream[j].Date)
	})

}
