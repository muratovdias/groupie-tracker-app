package handlers

import (
	"log"
	"net/http"
	"text/template"

	"tracker/internal/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Err(w, http.StatusMethodNotAllowed)
		log.Println("Search handler wrong method")
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	text := r.FormValue("search")
	if text == "" {
		Err(w, http.StatusBadRequest)
		log.Println("empty text value")
		return
	}
	artist, err := utils.SearchGroup(text)
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	err = templ.Execute(w, artist)
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}
