package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/routes"
	"my_podcast_api/util"
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
	config := &models.Config{}
	decoder.Decode(&config)

	conf := fmt.Sprintf("%s:%s@/%s", config.User, config.Password, config.Schema)
	db, err := gorm.Open("mysql", conf)

	if err != nil {
		log.Fatal(err)
	}

	jwtTokenUtil := &util.JwtTokenUtil{SigningKey: config.SigningKey}
	emailValidator := &validation.EmailValidation{}
	userDB := &repository.UserDB{db}
	episodeDB := &repository.EpisodeDB{db}
	podcastDB := &repository.PodcastDB{db}

	db.AutoMigrate(&models.User{}, &models.Podcast{}, &models.Episode{})

	defer userDB.Close()

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB})
	http.Handle("/createsession", &routes.CreateSessionHandler{DB: userDB, JwtTokenUtil: jwtTokenUtil})
	http.Handle("/getpodcasts", &routes.GetPodcastsHandler{UserDB: userDB, PodcastDB: podcastDB})
	http.Handle("/getepisodes", &routes.GetEpisodesHandler{UserDB: userDB, EpisodeDB: episodeDB})
	http.Handle("/upload", &routes.UploadEpisodeHandler{UserDB: userDB, EpisodeDB: episodeDB})

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
