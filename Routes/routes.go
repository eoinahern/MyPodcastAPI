package routes

import (
	"encoding/json"
	"io"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/util"
	"my_podcast_api/validation"
	"net/http"
)

//tregister user!!!
type RegisterHandler struct {
	EmailValidator *validation.EmailValidation
	DB             *repository.UserDB
}

type CreateSessionHandler struct {
	DB           *repository.UserDB
	JwtTokenUtil *util.JwtTokenUtil
}

type ReCreateSession struct {
}

type EndSessionHandler struct {
}

type GetPodcastsHandler struct {
	UserDB    *repository.UserDB
	PodcastDB *repository.PodcastDB
}

type GetEpisodesHandler struct {
	UserDB    *repository.UserDB
	EpisodeDB *repository.EpisodeDB
}

type DownloadEpisodeHandler struct {
	//need folder location!! key and credentials.
	//transmit file across network
}

type UploadEpisodeHandler struct {
	//credentials. then upload to network
	UserDB    *repository.UserDB
	EpisodeDB *repository.EpisodeDB
}

type DeleteEpisodeHandler struct {
}

func (r *RegisterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	if len(user.UserName) == 0 || len(user.Password) == 0 {
		http.Error(w, "incorrect params", http.StatusBadRequest)
		return
	}

	if isValidEmail := r.EmailValidator.CheckEmailValid(user.UserName); !isValidEmail {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	//check user is in the DB?

	if r.DB.CheckExist(user.UserName) {
		http.Error(w, "user already exists!!", http.StatusConflict)
		return
	}

	//encrypt password!! and then insert!!
	r.DB.Create(user)

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(user)
	w.Write(resp)
}

func (c *CreateSessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {

		decoder := json.NewDecoder(req.Body)
		var user models.User
		err := decoder.Decode(&user)

		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}

		if c.DB.ValidatePasswordAndUser(user.UserName, user.Password) {
			user.Token = c.JwtTokenUtil.CreateToken(user.UserName)
			jsonUser, _ := json.Marshal(user)
			w.Write(jsonUser)
		}
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (e *EndSessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "session ended!!!")

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

// get a list of the most popular podcasts and return to the users
// in json format

func (g *GetPodcastsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

//get episodes from a specific podcast

func (e *GetEpisodesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

func (e *UploadEpisodeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//need header authorisation. check legit?
	//if so add files info to database.
	//then upload file to folder. username/podcastname/files.extension

}
