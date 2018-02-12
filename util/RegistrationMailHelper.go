package util

import (
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

	auth := smtp.PlainAuth("", username, "hellothere123", host)

}

func (m *Mail) BuildMail() string {
	return ""
}

func (m *Mail) ConstructTemplate() {

}
