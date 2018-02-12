package util

type Mail struct {
	SenderId     string
	ToId         string
	Subject      string
	BodyLocation string
}

func (m *Mail) SendMail() {

	//setup
	//send a mail

}

func (m *Mail) BuildMail() string {
	return ""
}
