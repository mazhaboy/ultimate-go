package main

import (
	"net/http"

	"html/template"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var client *redis.Client
var templates *template.Template

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "index.html", comments)
}
