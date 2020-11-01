package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	githubUsername string = "lroolle"
)

type PageData struct {
	Host        string
	URL         string
	PageTitle   string
	UserName    string
	PackageName string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	packageName := strings.TrimLeft(strings.TrimSpace(r.URL.Path), "/")
	data := PageData{
		Host:        r.Host,
		URL:         r.URL.Path,
		PageTitle:   packageName,
		UserName:    githubUsername,
		PackageName: packageName,
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Executing index template:", err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.FileServer(http.Dir(".")))

	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Println("Server error:", err)
	}
}
