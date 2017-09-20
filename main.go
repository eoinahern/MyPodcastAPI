package main


import ( "net/http"
"my_podcast_api/routes")

func main() {
  http.Handle("/register", &RegisterHandler{})
  http.ListenAndServe(":8080", nil)
}