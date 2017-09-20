package routes

import (
	"fmt"
	"io"
	"net/http"
)

type RegisterHandler struct {
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

	if req.Method == http.MethodPost {
		fmt.Println("user is registered weeee!")
		return
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
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
