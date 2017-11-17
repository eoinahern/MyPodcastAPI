package routes

import (
	"encoding/json"
	"io"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/validation"
	"net/http"
)

//tregister user!!!
type RegisterHandler struct {
	EmailValidator *validation.EmailValidation
	DB             *repository.UserDB
}

type CreateSessionHandler struct {
}

type EndSessionHandler struct {
}

type GetPodcastsHandler struct {
}

type GetEpisodesHandler struct {
}

type DownloadEpisodeHandler struct {
}

type UploadEpisodeHandler struct {
}

type DeleteEpisodeHandler struct {
}

func (r *RegisterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// this is overly verbose and coupled need to refactor
	//1. check email is valid. and check email doesnt exist in DB.
	//2. send verification email!!!
	//3. return Gson User not verified!

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

	//username := req.Body("UserName")
	//password := req.Body("Password")

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

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(models.User{
		UserName: user.UserName,
		Verified: false,
		Password: user.Password,
	})
	w.Write(resp)
}

func (c *CreateSessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		pass := req.FormValue("password")
		if len(pass) > 0 {
			//and user name and pass existy in the DB then
			// return the key to be used by the user. create UUID!!
			io.WriteString(w, "10101010101010110")
			return
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
