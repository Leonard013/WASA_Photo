package database

import (
	"errors"
	"io/ioutil"
	"os"
	"time"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) GetStream(streamId string) ([]PhotoForStream, error) {
	if _, ok := images_stream[streamId]; !ok {

		rows, err := db.c.Query("SELECT p.photoId,p.title,p.photoPath,p.date,p.author FROM Photos p WHERE p.author IN (SELECT f.followedId FROM Follow f WHERE f.profileId = ? AND f.followedId NOT IN (SELECT b.bannerId FROM Ban b WHERE b.bannedId = ?))", streamId, streamId)
		if rows.Err() != nil {
			return nil, rows.Err()
		}

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

			t, err := time.Parse("2006-01-02 15:04:05", date)
			if !errors.Is(err, nil) {
				return nil, err
			}

			// Open the file
			file, err := os.Open(path)
			if !errors.Is(err, nil) {
				return nil, err
			}
			defer file.Close() // Make sure to close the file when done

			// Read the file into a byte slice ([]byte)
			data, err := ioutil.ReadAll(file)
			if !errors.Is(err, nil) {
				return nil, err
			}

			rows_2, err := db.c.Query("SELECT l.likeId, l.author FROM Likes l WHERE l.photoId = ? ", photoId)
			if rows.Err() != nil {
				return nil, rows.Err()
			}

			if !errors.Is(err, nil) {
				return nil, err
			}
			defer rows_2.Close()

			var likeIds []string
			var likeAuthors []string

			for rows_2.Next() {
				var likeId string
				var like_author string
				err = rows_2.Scan(&likeId, &like_author)
				if !errors.Is(err, nil) {
					return nil, err
				}
				likeIds = append(likeIds, likeId)
				likeAuthors = append(likeAuthors, like_author)
			}

			rows_3, err := db.c.Query("SELECT c.commentId, c.author, c.text, c.date FROM Comments c WHERE c.photoId = ? ", photoId)
			if rows.Err() != nil {
				return nil, rows.Err()
			}

			if !errors.Is(err, nil) {
				return nil, err
			}
			defer rows_3.Close()

			var commentIds []string
			var commentAuthors []string
			var commentTexts []string
			var commentDates []string

			for rows_3.Next() {
				var commentId string
				var comment_author string
				var comment_text string
				var comment_date string
				err = rows_3.Scan(&commentId, &comment_author, &comment_text, &comment_date)
				if !errors.Is(err, nil) {
					return nil, err
				}
				var comment_username string
				err = db.c.QueryRow("SELECT username FROM Users WHERE userId=?", comment_author).Scan(&comment_username)
				if !errors.Is(err, nil) {
					return nil, err
				}
				commentIds = append(commentIds, commentId)
				commentAuthors = append(commentAuthors, comment_username)
				commentTexts = append(commentTexts, comment_text)
				commentDates = append(commentDates, comment_date)
			}

			var username string
			err = db.c.QueryRow("SELECT username FROM Users WHERE userId=?", author).Scan(&username)
			if !errors.Is(err, nil) {
				return nil, err
			}

			photo := PhotoForStream{
				PhotoId:        photoId,
				Title:          title,
				File:           data,
				Author:         author,
				Username:       username,
				Date:           t,
				LikeIds:        likeIds,
				LikeAuthors:    likeAuthors,
				CommentIds:     commentIds,
				CommentAuthors: commentAuthors,
				CommentTexts:   commentTexts,
				CommentDates:   commentDates,
			}

			images_stream[streamId] = append(images_stream[streamId], photo)
		}
	}

	return images_stream[streamId], nil

}
