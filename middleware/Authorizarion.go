package middleware

import (
	"encoding/json"
	"fmt"
	"my_podcast_api/models"
	"my_podcast_api/util"
	"net/http"
	"strings"
)

type Authorization struct {
	Next         http.Handler
	JwtTokenUtil *util.JwtTokenUtil
}

func (a *Authorization) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	fmt.Println("authorization middleware!!")

	token := getTokenFromHeader(req)
	code, message := a.JwtTokenUtil.CheckTokenCredentials(token)

	if code != -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		msg, _ := json.Marshal(models.Message{Message: message})
		w.Write(msg)
		fmt.Println("auth failed!!")
		return
	}

	a.Next.ServeHTTP(w, req)
}

func getTokenFromHeader(req *http.Request) string {

	token := req.Header.Get("Authorization")
	tokenSlice := strings.Split(token, " ")

	if len(tokenSlice) != 2 {
		return ""
	}

	return tokenSlice[1]
}
