package main

import (
	"html/template"
	"log"
	"net/http"
)

func serveRoot(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func serveRootDeepti(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("indexdeepti.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/deepti/", serveRootDeepti)
	http.ListenAndServe(":8080", nil)
}