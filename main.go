package main


import ( "net/http"
"my_podcast_api/routes")

func main() {
  http.Handle("/register", &routes.RegisterHandler{})
  http.Handle("/createsession", &routes.CreateSessionHandler{})
  http.ListenAndServe(":8080", nil)
}