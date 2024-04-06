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

type UserId struct {
	UserId string `json:"userId"`
}

type phtoId_userId struct {
	PhotoId string `json:"photoId"`
	UserId  string `json:"userId"`
}

type User struct { // User represents a user
	UserId    string   `json:"userId"`
	Username  string   `json:"username"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Banned    []string `json:"banned"`
	IsBanned  []string `json:"isBanned"`
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
	PhotoId        string    `json:"photoId"`
	Title          string    `json:"title,omitempty"`
	File           []byte    `json:"File"`
	Author         string    `json:"author"`
	Username       string    `json:"username"`
	Date           time.Time `json:"date"`
	LikeIds        []string  `json:"likeIds"`
	LikeAuthors    []string  `json:"likeAuthors"`
	CommentIds     []string  `json:"commentIds"`
	CommentAuthors []string  `json:"commentAuthors"`
	CommentTexts   []string  `json:"commentTexts"`
	CommentDates   []string  `json:"commentDates"`
}

type Uncomment struct {
	PhotoId string `json:"photoId"`
	UserId  string `json:"userId"`
}

func SavePhoto(file multipart.File, id string) (string, error) {
	fileData, err := io.ReadAll(file)
	if err != nil {
		return "", errors.New("error reading file")
	}

	// Specify the directory to save the file
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	uploadPath := cwd + "/photos"
	//uploadPath := "/home/wasa/test-init/WASA/photos"

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
