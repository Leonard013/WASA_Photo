/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUsername(id string) (string, error)
	GetUserId(id string) (string, error)
	GetAuthorId(photoId string) (string, error)

	AddUser(username string) (string, error)
	UploadPhoto(photoId string, title string, photoPath string, userId string, time string) error

	SetName(name string) error

	CreateUserId() (string, error)
	IfBanned(id string, username string) error
	Follow(profileId string, followedId string) error
	Unfollow(profileId string, followedId string) error
	CheckToken(token string) error
	Ban(bannerId string, bannedId string) error
	Unban(bannerId string, bannedId string) error
	SetUsername(username string, userId string) error
	CreatePhotoId() (string, error)
	DeletePhoto(photoId string, userId string) (string, error)
	PutLike(photoId string, userId string) (string, error)
	CreateLikeId() (string, error)
	Unlike(photoId string, userId string) error
	Comment(photoId string, userId string, text string, t string) (string, error)
	Uncomment(commentId string, userId string) error
	GetStream(streamId string) ([]PhotoForStream, error)
	GetPhotos(userId string) ([]PhotoForStream, error)
	GetUserInfo(userId string) (UserInfo, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Create tables
	if err := CreateTables(db); !errors.Is(err, nil) {
		return nil, fmt.Errorf("error creating tables: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
