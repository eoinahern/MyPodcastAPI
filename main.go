package main

import (
	"encoding/json"
	"fmt"
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

	conf := fmt.Sprintf("%s:%s@/%s", dbConfig.User, dbConfig.Password, dbConfig.Schema)
	fmt.Println(conf)

	dbinst := &repository.DB{}
	db, err := dbinst.Open("mysql", conf)

	emailValidator := &validation.EmailValidation{}
	userDB := &repository.UserDB{db}
	episodeDB := &repository.EpisodeDB{db}
	podcastDB := &repository.PodcastDB{db}

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB})
	http.Handle("/createsession", &routes.CreateSessionHandler{})
	http.ListenAndServe(":8080", nil)
}
