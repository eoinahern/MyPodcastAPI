package models

type User struct {
	UserName string `json: "username"`
	Verified bool   `json: "verified"`
	Password string `json: "password"`
}

type Session struct {
	UserName   string
	SessionKey string
}

type Message struct {
	Message string `json: "message"`
}

type Podcast struct {
	Username string
	Icon     string
	Name     string //name of podcast
	Details  string //info about the podcast
}

type Episode struct {
	Episode   int32
	Created   string
	URL       string
	Downloads int32
	Details   string
}
