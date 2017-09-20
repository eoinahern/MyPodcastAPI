package main

import (
	"my_podcast_api/routes"
	"my_podcast_api/validation"
	"net/http"
)

func main() {

	emailValidator := &validation.EmailValidation{}

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator})
	http.Handle("/createsession", &routes.CreateSessionHandler{})
	http.ListenAndServe(":8080", nil)
}
