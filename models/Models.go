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
	Podcasts []Podcast `json: "podcasts" gorm: "ForeignKey:UserEmail;AssociationForeignKey:UserName"`
}

type Message struct {
	Message string `json: "message"`
}

type Podcast struct {
	UserEmail string    `json: "useremail" gorm: "type:TEXT"`
	Icon      string    `json: "icon"   gorm: "type:TEXT"`
	Name      string    `json: "name" gorm: "type: TEXT"`    //name of podcast
	Details   string    `json : "details" gorm: "type:TEXT"` //info about the podcast
	Episodes  []Episode `json: "episodes"  gorm: "ForeignKey:UserID; AssociationForeignKey:UserEmail"`
}

type Episode struct {
	UserID    string `json: "userid"  gorm: "type:Text"`
	Episode   int32  `json: "episode" gorm: "type: INTEGER"`
	Created   string `json: "created" gorm: "type: TEXT"`
	Updated   string `json: "updated" gorm: "type: TEXT"`
	URL       string `json: "url" gorm: "type: TEXT"`
	Downloads int32  `json: "downloads" gorm: "type: INTEGER"`
	Details   string `json: "details" gorm: "type: TEXT"`
}
