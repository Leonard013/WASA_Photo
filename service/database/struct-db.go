package database

import (
	"time"
)

var images_stream = make(map[string][]PhotoForStream)

type PhotoForStream struct { // PhotoForStream represents a photo for the stream
	PhotoId string    `json:"photoId"`
	Title   string    `json:"title,omitempty"`
	File    []byte    `json:"File"`
	Author  string    `json:"author"`
	Date    time.Time `json:"date"`
}


type UserInfo struct {
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Banned   []string `json:"banned"`
	IsBanned []string `json:"isBanned"`
}