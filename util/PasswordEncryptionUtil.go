package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptUtil struct {
}

func (p *PasswordEncryptUtil) Encrypt(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func (p *PasswordEncryptUtil) CheckSame(DBpassword string, sentPassword string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(DBpassword), []byte(sentPassword)); err != nil {
		return false
	} else {
		return true
	}
}
