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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	db, err := gorm.Open("mysql", conf)

	if err != nil {
		log.Fatal(err)
	}

	emailValidator := &validation.EmailValidation{}
	userDB := &repository.UserDB{db}
	episodeDB := &repository.EpisodeDB{db}
	podcastDB := &repository.PodcastDB{db}

	db.AutoMigrate(&models.User{}, &models.Podcast{}, &models.Episode{})

	defer userDB.Close()

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB})
	http.Handle("/createsession", &routes.CreateSessionHandler{DB: userDB})
	http.Handle("/getpodcasts", &routes.GetPodcastsHandler{UserDB: userDB, PodcastDB: podcastDB})
	http.Handle("/getepisodes", &routes.GetEpisodesHandler{UserDB: userDB, EpisodeDB: episodeDB})

	http.ListenAndServe(":8080", nil)
}
