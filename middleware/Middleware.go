package middleware

import (
	"encoding/json"
	"fmt"
	"my_podcast_api/models"
	"my_podcast_api/util"
	"net/http"
	"strings"
)

type Adapter func(http.Handler) http.Handler

func Adapt(finalHandler http.Handler, adapters ...Adapter) http.Handler {

	for _, item := range adapters {
		finalHandler = item(finalHandler)
	}

	return finalHandler
}

func AuthMiddlewareInit(jwtTokenUtil *util.JwtTokenUtil) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			fmt.Println("authorization middleware!!")
			token := getTokenFromHeader(req)
			code, message := jwtTokenUtil.CheckTokenCredentials(token)

			if code != -1 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(code)
				msg, _ := json.Marshal(models.Message{Message: message})
				w.Write(msg)
				fmt.Println("auth failed!!")
				return
			}

			h.ServeHTTP(w, req)

		})
	}
}

func getTokenFromHeader(req *http.Request) string {

	token := req.Header.Get("Authorization")
	tokenSlice := strings.Split(token, " ")

	if len(tokenSlice) != 2 {
		return ""
	}

	return tokenSlice[1]
}
