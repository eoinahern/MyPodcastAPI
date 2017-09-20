package models

type User struct {
	UserName string
	Password []byte
	Podcasts []Podcast
}

type Podcast struct {
	Username string
	Icon     string
	Name     string //name of podcast
	Details  string //info about the podcast
	Episodes []Episode
}

type Episode struct {
	Episode   int32
	Created   string
	URL       string
	Downloads int32
	Details   string
}
