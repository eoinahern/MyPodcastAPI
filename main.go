package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my_podcast_api/middleware"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/routes"
	"my_podcast_api/util"
	"my_podcast_api/validation"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	fileHelperUtil := &util.FileHelperUtil{}
	userDB := &repository.UserDB{db}
	episodeDB := &repository.EpisodeDB{db}
	podcastDB := &repository.PodcastDB{db}
	jwtTokenUtil := &util.JwtTokenUtil{SigningKey: config.SigningKey, DB: userDB}

	db.AutoMigrate(&models.User{}, &models.Podcast{}, &models.Episode{})
	db.Model(&models.Podcast{}).AddForeignKey("user_email", "users(user_name)", "CASCADE", "CASCADE")
	db.Model(&models.Episode{}).AddForeignKey("pod_id", "podcasts(podcast_id)", "CASCADE", "CASCADE")

	defer db.Close()

	router := mux.NewRouter()

	router.Handle("/register", &routes.RegisterHandler{EmailValidator: emailValidator, DB: userDB, PassEncryptUtil: passEncryptUtil})
	router.Handle("/createsession", &routes.CreateSessionHandler{DB: userDB, JwtTokenUtil: jwtTokenUtil, PassEncryptUtil: passEncryptUtil})
	router.Handle("/getpodcasts", &routes.GetPodcastsHandler{UserDB: userDB, PodcastDB: podcastDB, JwtTokenUtil: jwtTokenUtil})
	router.Handle("/getepisodes", &routes.GetEpisodesHandler{UserDB: userDB, EpisodeDB: episodeDB, JwtTokenUtil: jwtTokenUtil})
	router.Handle("/download/{podcastid}/{podcastname}/{podcastfilename}", middleware.Adapt(&routes.DownloadEpisodeHandler{EpisodeDB: episodeDB}, middleware.StringMiddlewareInit("hello there"), middleware.AuthMiddlewareInit(jwtTokenUtil))).Methods(http.MethodGet) //&middleware.Authorization{JwtTokenUtil: jwtTokenUtil, Next: &routes.DownloadEpisodeHandler{EpisodeDB: episodeDB}}).Methods(http.MethodGet)
	router.Handle("/createpodcast", &middleware.Authorization{JwtTokenUtil: jwtTokenUtil, Next: &routes.CreatePodcastHandler{PodcastDB: podcastDB, FileHelper: fileHelperUtil}}).Methods(http.MethodPost)
	router.Handle("/upload", &middleware.Authorization{JwtTokenUtil: jwtTokenUtil, Next: &routes.UploadEpisodeHandler{UserDB: userDB, PodcastDB: podcastDB, EpisodeDB: episodeDB}}).Methods(http.MethodPost)

	http.ListenAndServe(":8080", router)
	//http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
