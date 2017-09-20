package routes

import (
	"encoding/json"
	"io"
	"my_podcast_api/models"
	"my_podcast_api/validation"
	"net/http"
)

type RegisterHandler struct {
	EmailValidator *validation.EmailValidation
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

	//1. check email is valid. and check email doesnt exist in DB.
	//2. check password length complexity
	//3. if both valid return "200" user created!!

	//4. return json with error message

	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := req.FormValue("UserName")
	password := req.FormValue("Password")

	if len(username) == 0 || len(password) == 0 {
		http.Error(w, "incorrect params", http.StatusBadRequest)
		return
	}

	if isValidEmail := r.EmailValidator.CheckEmailValid(username); !isValidEmail {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(models.Message{"it worked!!!"})
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
