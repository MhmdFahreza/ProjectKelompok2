package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

var films = []Film{
	{Title: "The Godfather", Director: "Francis Ford Coppola"},
	{Title: "Blade Runner", Director: "Ridley Scott"},
	{Title: "The Thing", Director: "John Carpenter"},
}

func main() {
	fmt.Println("Running Web...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, films)
	}

	// handler function #2 - returns the template block with the newly added film, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		films = append(films, Film{Title: title, Director: director})
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list", films)
	}

	// Serve static files (CSS and JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
