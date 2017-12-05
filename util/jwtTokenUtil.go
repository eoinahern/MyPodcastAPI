package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtTokenUtil struct {
	SigningKey string
}

func (j *JwtTokenUtil) CreateToken(username string) string {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour + 1).Unix()

	signedToken, _ := token.SignedString(j.signingKey)

	return signedToken

}

//make sure the token sent is correct!!!

func (j *JwtTokenUtil) CheckTokenCredentials(token *jwt.Token) bool {

}
