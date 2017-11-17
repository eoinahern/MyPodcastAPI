package main

import (
	"my_podcast_api/repository"
	"my_podcast_api/routes"
	"my_podcast_api/validation"
	"net/http"
)

func main() {

	emailValidator := &validation.EmailValidation{}
	userDB := &repository.UserDB{}

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB})
	http.Handle("/createsession", &routes.CreateSessionHandler{})
	http.ListenAndServe(":8080", nil)
}
