package database

import (
	"time"
)

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

type UserInfo struct {
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Banned    []string `json:"banned"`
	IsBanned  []string `json:"isBanned"`
}
