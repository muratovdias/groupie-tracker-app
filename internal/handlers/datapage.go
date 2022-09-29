package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"tracker/internal/utils"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artists/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Err(w, http.StatusNotFound)
		log.Println("ArtistPage, Atoi function")
		return
	}
	if r.Method != http.MethodGet {
		Err(w, http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("templates/about.html")
	if err != nil {
		log.Println("artist page ParseFiles function", err)
		Err(w, http.StatusInternalServerError)
		return
	}
	artists, err := utils.GetData()
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	if id < 0 || id > len(artists.Artists) {
		Err(w, http.StatusNotFound)
		return
	}
	err = templ.Execute(w, &artists.Artists[id-1])
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}
