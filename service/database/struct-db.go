package database

import "time"

var Users = make(map[string]User)

var streams []stream_reminder

type stream_reminder struct {
	UserId         string   `json:"userId"`
	Counter        int      `json:"counter"`
	List_of_Photos [][]byte `json:"list_of_photos"`
}

type Photo struct {
	Image []byte
	Date  time.Time
}
