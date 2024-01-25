package api

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"time"
)

var Users = make(map[string]User)              // Users is a map of users
var UsersFollowers = make(map[string][]string) // UsersFollowers is a map of users' followers
var UsersFollowing = make(map[string][]string) // UsersFollowing is a map of users' following
var Photos = make(map[string]Photo)            // Photos is a map of photos
var Comments = make(map[string]Comment)        // Comments is a map of comments
var Likes = make(map[string]Like)              // Likes is a map of likes
var images_stream = make(map[string][]PhotoForStream)

var UsersCatalog = make(map[string]string) // UsersCatalog is a map of users' usernames and ids

type Username struct {
	Username string `json:"username"`
}
type phtoId_userId struct {
	PhotoId string `json:"photoId"`
	UserId  string `json:"userId"`
}
type User_Id struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type User struct { // User represents a user
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

type Photo struct { // Photo represents a photo
	PhotoId   string `json:"photoId"`
	Title     string `json:"title,omitempty"`
	PhotoPath string `json:"photoPath"`
	Date      string `json:"date"`
	Author    string `json:"userId"`
}

type Comment struct { // Comment represents a comment
	PhotoId   string `json:"photoId"`
	CommentId string `json:"commentId"`
	Author    string `json:"userId"`
	Text      string `json:"text"`
	Date      string `json:"date"`
}

type Like struct { // Like represents a like
	LikeId  string `json:"likeId"`
	PhotoId string `json:"photoId"`
	Author  string `json:"userId"`
}

type PhotoForStream struct { // PhotoForStream represents a photo for the stream
	PhotoId string    `json:"photoId"`
	Title   string    `json:"title,omitempty"`
	File    string    `json:"File"`
	Author  string    `json:"author"`
	Date    time.Time `json:"date"`
}

func SavePhoto(file multipart.File, id string) (string, error) {
	fileData, err := io.ReadAll(file)
	if err != nil {
		return "", errors.New("error reading file")
	}

	// Specify the directory to save the file
	uploadPath := "/Users/leonardoscappatura/Documents/GitHub/WASA/photos"

	// Create/Open the file
	filePath := uploadPath + "/" + id + ".png"
	newFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	// Write the file data
	_, err = newFile.Write(fileData)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func DeletePhoto(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
