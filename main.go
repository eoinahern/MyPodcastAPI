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
		fmt.Println(err)
		log.Fatal(err)
	}

	//create dependencies
	passEncryptUtil := &util.PasswordEncryptUtil{}
	emailValidator := &validation.EmailValidation{}
	fileHelperutil := &util.FileHelperUtil{}
	userDB := &repository.UserDB{db}
	episodeDB := &repository.EpisodeDB{db}
	podcastDB := &repository.PodcastDB{db}
	jwtTokenUtil := &util.JwtTokenUtil{SigningKey: config.SigningKey, DB: userDB}

	db.AutoMigrate(&models.User{}, &models.Podcast{}, &models.Episode{})
	db.Model(&models.Podcast{}).AddForeignKey("user_email", "users(user_name)", "CASCADE", "CASCADE")
	db.Model(&models.Episode{}).AddForeignKey("pod_id", "podcasts(podcast_id)", "CASCADE", "CASCADE")

	defer db.Close()

	http.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB, PassEncryptUtil: passEncryptUtil})
	http.Handle("/createsession", &routes.CreateSessionHandler{DB: userDB, JwtTokenUtil: jwtTokenUtil, PassEncryptUtil: passEncryptUtil})
	http.Handle("/getpodcasts", &routes.GetPodcastsHandler{UserDB: userDB, PodcastDB: podcastDB, JwtTokenUtil: jwtTokenUtil})
	http.Handle("/getepisodes", &routes.GetEpisodesHandler{UserDB: userDB, EpisodeDB: episodeDB})
	http.Handle("/createpodcast", &routes.CreatePodcastHandler{PodcastDB: podcastDB, JwtTokenUtil: jwtTokenUtil, FileHelper: fileHelperutil})
	http.Handle("/upload", &routes.UploadEpisodeHandler{UserDB: userDB, EpisodeDB: episodeDB, JwtTokenUtil: jwtTokenUtil})

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
