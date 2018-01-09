package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Port       string `json: "port"`
	Password   string `json: "password"`
	User       string `json: "user"`
	Schema     string `json: "schema"`
	SigningKey string `json: "signingkey"`
}

type User struct {
	UserName string    `json: "username" gorm: "primary_key; type:TEXT; not null; unique"`
	Verified bool      `json: "verified" gorm: "type : "BOOLEAN" `
	Password string    `json: "password" gorm: "type: TEXT"`
	Token    string    `json: "token" sql:"-" gorm:"-" `
	Podcasts []Podcast `json: "podcasts" gorm: "ForeignKey:UserEmail"`
}

type UserTitle struct {
	UserName string `json: "username"`
}

type Message struct {
	Message string `json: "message"`
}

type Podcast struct {
	PodcastID  uint      `gorm: "primary_key; AUTO_INCREMENT; type:INTEGER; not null" json: "podcastid"`
	UserEmail  string    `json: "useremail" gorm: "type:TEXT"`
	Icon       string    `json: "icon"   gorm: "type:TEXT"`
	Name       string    `json: "name" gorm: "type: TEXT"`
	EpisodeNum int       `json: "episodenum" gorm: "type:INTEGER; default:0"`
	Details    string    `json : "details" gorm: "type:TEXT"`
	Episodes   []Episode `json: "episodes"  gorm: "ForeignKey:PodID"`
}

type Episode struct {
	PodID     uint   `gorm: "type:INTEGER" json: "podid"`
	Created   string `json: "created" gorm: "type: TEXT"`
	Updated   string `json: "updated" gorm: "type: TEXT"`
	URL       string `json: "url" gorm: "type: TEXT"`
	Downloads int32  `json: "downloads" gorm: "type: INTEGER; not null default:0"`
	Blurb     string `json: "blurb" gorm: "type: TEXT"`
}
