package handlers

import (
	"net/http"
	"text/template"
)

func Err(w http.ResponseWriter, status int) {
	Err := ErrorHTTP{status, http.StatusText(status)}
	templ, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(Err.Status)
	err = templ.Execute(w, Err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type ErrorHTTP struct {
	Status  int
	Message string
}
