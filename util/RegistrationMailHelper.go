package util

const from string = "mypodcastapi@gmail.com"

type Mail struct {
	SenderId string
	ToId     string
	Subject  string
	Body     *[]byte
}

func (m *Mail) SendMail() {

	//setup
	//send a mail

}

func (m *Mail) CreateMessageTemplate(file string) []byte {
	return []byte{}
}
