package routes

import ( "net/http"
 "io"
"fmt")



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
  
  if req.Method == http.MethodPost { 
    io.WriteString(w, "user is registered")
    fmt.Println("user is registered weeee!")
    return
  } else {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
  }
}

  
func (c *CreateSessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  if req.Method == http.MethodPost { 
     
     pass :=  req.FormValue("password")
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

