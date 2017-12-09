package util

import (
	"log"
	"my_podcast_api/repository"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtTokenUtil struct {
	SigningKey string
	DB         *repository.UserDB
}

func (j *JwtTokenUtil) CreateToken(username string) string {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour + 1).Unix()

	signedToken, err := token.SignedString([]byte(j.SigningKey))

	if err != nil {
		log.Fatal(err)
	}

	return signedToken
}

func (j *JwtTokenUtil) CheckTokenCredentials(tokenStr string, userName string) (int, string) {

	token, err := jwt.Parse(tokenStr, func(passedToken *jwt.Token) (interface{}, error) {
		return []byte(j.SigningKey), nil
	})

	if err != nil {
		return http.StatusUnauthorized, "error validating token"
	}

	claims := token.Claims.(jwt.MapClaims)
	time := int64(claims["exp"].(float64))
	name := claims["name"].(string)

	if !verifyTokenTime(time) || !j.verifyTokenUser(name, userName) {
		return http.StatusUnauthorized, "error validating token"
	}

	return -1, ""
}

func verifyTokenTime(chimey int64) bool {
	if chimey > time.Now().Unix() {
		log.Println("token isnt expired")
		return true
	} else {
		log.Println("token is expired")
		return false
	}
}

func (j *JwtTokenUtil) verifyTokenUser(tokenName string, userName string) bool {

	if strings.Compare(tokenName, userName) != 0 {
		log.Println("name comparison failed")
		return false
	}

	if !j.DB.CheckExist(userName) {
		return false
	}

	log.Println("name comparison succeed")
	return true
}
