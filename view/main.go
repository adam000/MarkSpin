package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/adam000/goutils/page"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func Main() {
	tpl = template.New("T")
	template.Must(tpl.ParseGlob("templates/main/*"))
	template.Must(tpl.ParseGlob("web-resources/templates/*"))
	handle()
}

func handle() {
	page.SetSiteTitle("MarkSpin")
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/_ah/health", healthCheckHandler)
	/*
		TODO
			r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "static/favicon.ico")
			})
	*/
	r.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/robots.txt")
	})

	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	initLetsEncryptHandler(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

type project struct {
	Name        string
	Link        string
	Description string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var page = page.NewPage()
	page.SetTitle("Hello World")
	tpl.ExecuteTemplate(w, "page_home.html", page)
}

func initLetsEncryptHandler(r *mux.Router) {
	r.HandleFunc("/.well-known/acme-challenge/{stuff}", challengeHandler)
}

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/acme/challenge")
}
