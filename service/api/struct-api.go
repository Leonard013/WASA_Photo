package api

var Users = make(map[string]User)              // Users is a map of users
var UsersFollowers = make(map[string][]string) // UsersFollowers is a map of users' followers
var UsersFollowing = make(map[string][]string) // UsersFollowing is a map of users' following
// var bannedUsers = make(map[string][]string)    // bannedUsers is a map of banned users
var Photos = make(map[string]Photo)     // Photos is a map of photos
var Comments = make(map[string]Comment) // Comments is a map of comments
var Likes = make(map[string]Like)       // Likes is a map of likes

var UsersCatalog = make(map[string]string) // UsersCatalog is a map of users' usernames and ids

type User struct { // User represents a user
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

type Photo struct { // Photo represents a photo
	PhotoId   string `json:"photoId"`
	Title     string `json:"title,omitempty"`
	PhotoPath string `json:"photoPath"`
	Date      string `json:"date"`
	UserId    string `json:"userId"`
}

type Comment struct { // Comment represents a comment
	PhotoId   string `json:"photoId"`
	CommentId string `json:"commentId"`
	UserId    string `json:"userId"`
	Text      string `json:"text"`
	Date      string `json:"date"`
}

type Like struct { // Like represents a like
	LikeId  string `json:"likeId"`
	PhotoId string `json:"photoId"`
	UserId  string `json:"userId"`
}