package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptUtil struct {
}

func (p *PasswordEncryptUtil) encrypt(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return hash
}

func (p *PasswordEncryptUtil) checkSame(DBpassword string, sentPassword string) bool {

	if err := bcrypt.CompareHashAndPassword(DBpassword, []byte(sentPassword)); err != nil {
		return false
	} else {
		return true
	}
}
