package database

import (
	"errors"
	"os"
	"time"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) GetStream(userId string) ([]PhotoForStream, error) {
	if _, ok := images_stream[userId]; !ok {

		rows, err := db.c.Query("SELECT p.photoId,p.title,p.photoPath,p.date,p.author FROM Photos p WHERE p.author IN (SELECT f.followedId FROM Follow f WHERE f.profileId = ? AND f.followedId NOT IN (SELECT b.bannerId FROM Ban b WHERE b.bannedId = ?))", userId, userId)
		if !errors.Is(err, nil) {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {

			var path string
			var title string
			var date string
			var photoId string
			var author string
			err = rows.Scan(&photoId, &title, &path, &date, &author)
			if !errors.Is(err, nil) {
				return nil, err
			}

			img, err := os.Open(path)
			if !errors.Is(err, nil) {
				return nil, err
			}
			defer img.Close()

			t, err := time.Parse("2006-01-02 15:04:05", date)
			if !errors.Is(err, nil) {
				return nil, err
			}

			photo := PhotoForStream{
				PhotoId: photoId,
				Title:   title,
				File:    img,
				Author:  author,
				Date:    t,
			}

			images_stream[userId] = append(images_stream[userId], photo)
		}
	}

	return images_stream[userId], nil

}