package util

import (
	"net/smtp"
)

const (
	from    string = "mypodcastapi@gmail.com"
	message string = "hello, you are now registere as a user of the app blah" // want to make this a styled email
)

func SendMail(userEmail *string) {

 err := smtp.SendMail(addr, , from, *userEmail, message)

 if err != nil {
	 log.Println(err)
 }


}
