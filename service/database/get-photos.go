package database

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) GetPhotos(userId string) ([]PhotoForStream, error) {
	var photos []PhotoForStream
	rows, err := db.c.Query("SELECT p.photoId,p.title,p.photoPath,p.date,p.author FROM Photos p WHERE p.author = ", userId)
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

		photos = append(photos, photo)
	}
	fmt.Println(photos)

	return photos, nil

}
