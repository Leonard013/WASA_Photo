package database

import (
	"database/sql"
	"fmt"
)

// CreateTables initializes the database tables based on the OpenAPI specification.
func CreateTables(db *sql.DB) error {
	// Create User table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			userId TEXT NOT NULL PRIMARY KEY,
			username TEXT NOT NULL,
			UNIQUE(username)
			UNIQUE(userId)
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating User table: %w", err)
	}

	// Create Photo table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Photos (
			photoId TEXT NOT NULL PRIMARY KEY,
			title TEXT,
			photoPath TEXT NOT NULL,
			date TEXT NOT NULL,
			author TEXT NOT NULL,
			FOREIGN KEY (author) REFERENCES User(userId) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating Photo table: %w", err)
	}

	// Create Like table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Likes (
			likeId TEXT NOT NULL PRIMARY KEY,
			photoId TEXT NOT NULL,
			author TEXT NOT NULL,
			FOREIGN KEY (photoId) REFERENCES Photo(photoId) ON DELETE CASCADE,
			FOREIGN KEY (author) REFERENCES User(userId) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating Like table: %w", err)
	}

	// Create Comment table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Comments (
			commentId TEXT NOT NULL PRIMARY KEY,
			photoId TEXT NOT NULL,
			author TEXT NOT NULL,
			text TEXT,
			date TEXT,
			FOREIGN KEY (photoId) REFERENCES Photo(photoId) ON DELETE CASCADE,
			FOREIGN KEY (author) REFERENCES User(userId) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating Comment table: %w", err)
	}

	// Create Following table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Follow (
			profileId TEXT NOT NULL,
			followedId TEXT NOT NULL,
			PRIMARY KEY (profileId, followedId),
			FOREIGN KEY (profileId) REFERENCES User(userId) ON DELETE CASCADE,
			FOREIGN KEY (followedId) REFERENCES User(userId) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating Following table: %w", err)
	}

	// Create Banned table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Ban (
			bannerId TEXT NOT NULL,
			bannedId TEXT NOT NULL,
			PRIMARY KEY (bannedId, bannerId),
			FOREIGN KEY (bannedId) REFERENCES User(userId) ON DELETE CASCADE,
			FOREIGN KEY (bannerId) REFERENCES User(userId) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating Banned table: %w", err)
	}

	return nil
}
