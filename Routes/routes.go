package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/util"
	"my_podcast_api/validation"
	"net/http"
	"strings"
)

//tregister user!!!
type RegisterHandler struct {
	EmailValidator  *validation.EmailValidation
	DB              *repository.UserDB
	PassEncryptUtil *util.PasswordEncryptUtil
}

type CreateSessionHandler struct {
	DB              *repository.UserDB
	JwtTokenUtil    *util.JwtTokenUtil
	PassEncryptUtil *util.PasswordEncryptUtil
}

type ReCreateSession struct {
}

type EndSessionHandler struct {
}

type CreatePodcastHandler struct {
	PodcastDB    *repository.PodcastDB
	JwtTokenUtil *util.JwtTokenUtil
	FileHelper   *util.FileHelperUtil
}

type GetPodcastsHandler struct {
	UserDB       *repository.UserDB
	PodcastDB    *repository.PodcastDB
	JwtTokenUtil *util.JwtTokenUtil
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
	UserDB       *repository.UserDB
	EpisodeDB    *repository.EpisodeDB
	PodcastDB    *repository.PodcastDB
	JwtTokenUtil *util.JwtTokenUtil
}

type DeleteEpisodeHandler struct {
	UserDB    *repository.UserDB
	PodcastDB *repository.PodcastDB
}

//vars
var tokenErr []byte = []byte(`{ "error" : "problem with token"}`)
var internalErr []byte = []byte(`{ "error" : "internal error"}`)

const notAllowedErrStr string = "method not allowed"
const podcastFiles string = "../files"

/**
*	helper to get auth token
*
**/

func getTokenFromHeader(req *http.Request) string {

	token := req.Header.Get("Authorization")
	tokenSlice := strings.Split(token, " ")

	if len(tokenSlice) != 2 {
		return ""
	}

	return tokenSlice[1]
}

func (r *RegisterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, notAllowedErrStr, http.StatusMethodNotAllowed)
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

	user.Password = r.PassEncryptUtil.Encrypt(user.Password)
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

		dbUser := c.DB.GetUser(user.UserName)
		w.Header().Set("Content-Type", "application/json")

		if c.PassEncryptUtil.CheckSame(dbUser.Password, user.Password) {
			user.Token = c.JwtTokenUtil.CreateToken(user.UserName)
			jsonUser, _ := json.Marshal(user)
			w.Write(jsonUser)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error" : "incorrect pass"}`))
		}

	} else {
		http.Error(w, notAllowedErrStr, http.StatusMethodNotAllowed)
	}

}

func (e *EndSessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "session ended!!!")

	} else {
		http.Error(w, notAllowedErrStr, http.StatusMethodNotAllowed)
	}

}

//create a podcast enntry and folder on server.

func (c *CreatePodcastHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//1. authorize user .....
	//2. create user folder if not exists. create podcast folder.
	//3. store data in DB about podcast
	//3. return success

	if req.Method == http.MethodPost {

		w.Header().Set("Content-Type", "application/json")

		token := getTokenFromHeader(req)
		code, _ := c.JwtTokenUtil.CheckTokenCredentials(token)

		if code != -1 {
			w.WriteHeader(code)
			w.Write(tokenErr)
			return
		}

		//navigate to use dir!!!

		podcastname := req.URL.Query().Get("name")

		if len(podcastname) == 0 {
			http.Error(w, http.StatusText(22), http.StatusBadRequest)
			return
		}

		path := fmt.Sprintf("%s/%s/%s", podcastFiles, "myname", podcastname)
		fmt.Println(path)

		if !c.FileHelper.CheckDirFileExists(path) {
			c.FileHelper.CreateDir(path)
		}

		//add data to DB return success

	} else {
		http.Error(w, notAllowedErrStr, http.StatusMethodNotAllowed)
	}

}

// get a list of the most popular podcasts and return to the users
// in json format

func (g *GetPodcastsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//1. authorize user...
	//2. if authenticated. return most popular podcasts based on num downloads

	if req.Method == http.MethodGet {

		token := getTokenFromHeader(req)
		w.Header().Set("Content-Type", "application/json")
		code, _ := g.JwtTokenUtil.CheckTokenCredentials(token)

		if code != -1 {
			w.WriteHeader(code)
			w.Write(tokenErr)
			return
		}

		podcasts := g.PodcastDB.GetAll()
		podcastsMarshaled, err := json.Marshal(podcasts)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(internalErr)
		} else {
			w.Write(podcastsMarshaled)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{ "error" : "not allowed"}`))
	}

}

/**
*	get episodes from a specific podcast
* by podcast id and name!!!
**/

func (e *GetEpisodesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

func (e *UploadEpisodeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//need to pass podcast obj with slice of episodes

	//1.need header authorisation. check legit?
	//2.if so add files info to database.
	//3.then upload file to folder. username/podcastname/files.extension

	//decode obj sent
	decoder := json.NewDecoder(req.Body)
	var podcast models.Podcast
	err := decoder.Decode(&podcast)

	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	token := getTokenFromHeader(req)
	req.Header.Get("name")
	code, message := e.JwtTokenUtil.CheckTokenCredentials(token)

	if code != -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		msg, _ := json.Marshal(models.Message{Message: message})
		w.Write(msg)
	}

	//1. check podcast name? if it doesnt exist create? file extension check?
	//2. check podcast number from podcasts table. increment count by 1.
	//3. hash file name. read file and save in directory.
	//4. add file details to DB. return 200 ok

}
