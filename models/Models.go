package models

type DBConfig struct {
	Port     string `json: "port"`
	Password string `json: "password"`
	User     string `json: "user"`
	Schema   string `json: "schema"`
}

type User struct {
	UserName string `json: "username"  gorm: "type:TEXT; primary_key; not null; unique"`
	Verified bool   `json: "verified" gorm: "type : "BOOLEAN" `
	Password string `json: "password" gorm "type: TEXT"`
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
