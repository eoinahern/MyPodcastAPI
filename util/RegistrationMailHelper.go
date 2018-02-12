package util

import (
	"my_podcast_api/models"
	"net/smtp"
)

const host string = "smtp.gmail.com"

type Mail struct {
	SenderId     string
	ToId         string
	Subject      string
	BodyLocation string
	BodyParams   map[string]string
}

func (m *Mail) SendMail() {

	//create auth
	smtpConf := models.SmtpConfig{}
	smtpConf.ReadFromFile("../smtpConfig.json")
	auth := smtp.PlainAuth("", smtpConf.Username, smtpConf.Password, smtpConf.Server)

}

func (m *Mail) BuildMail() string {
	return ""
}

func (m *Mail) ConstructTemplate() {

}
