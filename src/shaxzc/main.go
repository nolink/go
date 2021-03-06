package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("about.html")
	t.Execute(w, nil)
}

func benefit(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("benefit.html")
	t.Execute(w, nil)
}

func contrast(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("contrast.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/benefit", benefit)
	http.HandleFunc("/contrast", contrast)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
