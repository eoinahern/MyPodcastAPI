package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"my_podcast_api/models"
	"my_podcast_api/repository"
	"my_podcast_api/util"
	"my_podcast_api/validation"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

//all episodes associated with specific podcast
type GetEpisodesHandler struct {
	UserDB       *repository.UserDB
	EpisodeDB    *repository.EpisodeDB
	JwtTokenUtil *util.JwtTokenUtil
}

//a specific episode data
type DownloadEpisodeHandler struct {
	EpisodeDB *repository.EpisodeDB
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
const podcastFiles string = "./files"

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
		http.Error(w, http.StatusText(31), http.StatusConflict)
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

//create a podcast entry and folder on server.

func (c *CreatePodcastHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//1. authorize user .....
	//2. create user folder if not exists. create podcast folder.
	//3. store data in DB about podcast
	//3. return success

	if req.Method == http.MethodPost {

		w.Header().Set("Content-Type", "application/json")

		var podcast models.Podcast
		err := json.NewDecoder(req.Body).Decode(&podcast)

		if err != nil {
			http.Error(w, http.StatusText(51), http.StatusInternalServerError)
			return
		}

		token := getTokenFromHeader(req)
		code, _ := c.JwtTokenUtil.CheckTokenCredentials(token)

		if code != -1 {
			w.WriteHeader(code)
			w.Write(tokenErr)
			return
		}

		podcastname := req.URL.Query().Get("podcastname")

		if len(podcastname) == 0 {
			http.Error(w, http.StatusText(22), http.StatusBadRequest)
			return
		}

		path := fmt.Sprintf("%s/%d/%s", podcastFiles, podcast.PodcastID, podcastname)

		if !c.FileHelper.CheckDirFileExists(path) {
			c.FileHelper.CreateDir(path)
			podcast.Location = path
			podcast.Name = podcastname
			err = c.PodcastDB.CreatePodcast(podcast)

			if err != nil {
				http.Error(w, http.StatusText(51), http.StatusInternalServerError)
				return
			}

			mpod, _ := json.Marshal(podcast)
			w.Write(mpod)
		}

	} else {
		http.Error(w, notAllowedErrStr, http.StatusMethodNotAllowed)
	}

}

// get a list of the most popular podcasts and return to the users
// in json format

func (g *GetPodcastsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

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
* by podcast id!!!
**/

func (g *GetEpisodesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {

		token := getTokenFromHeader(req)
		status, errTxt := g.JwtTokenUtil.CheckTokenCredentials(token)

		if status != -1 {
			http.Error(w, errTxt, status)
			return
		}

		podcastid, err := strconv.Atoi(req.URL.Query().Get("podcastid"))

		if err != nil {
			http.Error(w, http.StatusText(22), http.StatusBadRequest)
			return
		}

		episodes := g.EpisodeDB.GetAllEpisodes(podcastid)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&episodes)

	} else {
		http.Error(w, notAllowedErrStr, http.StatusMethodNotAllowed)
	}

}

func (g *DownloadEpisodeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	requestParams := mux.Vars(req)

	podcastID := requestParams["podcastid"]
	podcastName := requestParams["podcastname"]
	podcastFileName := requestParams["podcastfilename"]

	if len(podcastID) == 0 || len(podcastName) == 0 || len(podcastFileName) == 0 {
		http.Error(w, "unrecognised", http.StatusBadRequest)
		return
	}

	podlocation := fmt.Sprintf("%s/%s/%s/%s", podcastFiles, podcastID, podcastName, podcastFileName)
	fmt.Println(podlocation)
	filedata, err := ioutil.ReadFile(podlocation)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Write(filedata)

}

func (e *UploadEpisodeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var episode models.Episode
	file, fh, fileErr := req.FormFile("namefile")
	sepisode := req.FormValue("data")
	podcastname := req.URL.Query().Get("podcast")
	err := json.Unmarshal([]byte(sepisode), &episode)

	if len(sepisode) == 0 || err != nil || fileErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	podcast := e.PodcastDB.CheckPodcastCreated(episode.PodID, podcastname)

	if len(podcast.Name) == 0 {
		http.Error(w, "unknown podcast", http.StatusInternalServerError)
		return
	}

	splitname := strings.Split(fh.Filename, ".")
	ext := splitname[len(splitname)-1]

	if strings.Compare(ext, "mp3") != 0 {
		http.Error(w, "wrong file type", http.StatusInternalServerError)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	lastepisode := e.EpisodeDB.GetLastEpisode()
	filelocation := fmt.Sprintf("%s/%d.%s", podcast.Location, lastepisode.EpisodeID+1, "mp3")
	episode.URL = filelocation
	e.EpisodeDB.AddEpisode(episode)
	ioutil.WriteFile(filelocation, fileBytes, os.ModePerm)
	e.PodcastDB.UpdatePodcastNumEpisodes(podcast.PodcastID)

}
