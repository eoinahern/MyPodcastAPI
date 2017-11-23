package main

import (
	"encoding/json"
	"log"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/routes"
	"my_podcast_api/validation"
	"net/http"
	"os"
)

func main() {

	file, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	dbConfig := &models.DBConfig{}
	decoder.Decode(&dbConfig)

	emailValidator := &validation.EmailValidation{}
	userDB := &repository.UserDB{}

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB})
	http.Handle("/createsession", &routes.CreateSessionHandler{})
	http.ListenAndServe(":8080", nil)
}
