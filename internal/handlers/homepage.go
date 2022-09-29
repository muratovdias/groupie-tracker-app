package handlers

import (
	"log"
	"net/http"
	"text/template"

	"tracker/internal/utils"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Err(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Err(w, http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		log.Println("home page", err)
		return
	}

	artists, errs := utils.GetData()
	if errs != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	artists.FoundArtists = artists.Artists
	err = templ.Execute(w, &artists)
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}
