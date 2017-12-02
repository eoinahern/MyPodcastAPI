package models

type DBConfig struct {
	Port     string `json: "port"`
	Password string `json: "password"`
	User     string `json: "user"`
	Schema   string `json: "schema"`
}

type User struct {
	UserName string    `json: "username" gorm: "type:TEXT; primary_key; not null; unique"`
	Verified bool      `json: "verified" gorm: "type : "BOOLEAN" `
	Password string    `json: "password" gorm: "type: TEXT"`
	Podcasts []Podcast `json: "podcasts" gorm: "ForeignKey:UserName"`
}

type Message struct {
	Message string `json: "message"`
}

type Podcast struct {
	Icon     string    `json: "icon"`
	Name     string    `json: "name"`     //name of podcast
	Details  string    `json : "details"` //info about the podcast
	Episodes []Episode `json: "episodes"`
}

type Episode struct {
	Episode   int32  `json: "episode"`
	Created   string `json: "created"`
	Updated   string `json: "updated"`
	URL       string `json: "url"`
	Downloads int32  `json: downloads`
	Details   string `json: "details"`
}
