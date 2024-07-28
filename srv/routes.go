package main

import (
	"log"
	"net/http"
	"text/template"
)

func (app *application) getRoot(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/root.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
func (app *application) getbox(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/search.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
func (app *application) getlink(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/link.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
