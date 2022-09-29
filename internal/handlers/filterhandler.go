package handlers

import (
	"net/http"
	"text/template"

	"tracker/internal/utils"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter/" {
		Err(w, http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet {
		Err(w, http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	mincreationdate := r.FormValue("mindate")
	maxcreationdate := r.FormValue("maxdate")
	minfirstalbum := r.FormValue("minalbumdate")
	maxfirstalbum := r.FormValue("maxalbumdate")
	location := r.FormValue("location")
	r.ParseForm()
	members := r.Form["memberscount"]

	artist, err := utils.Filter_artists(mincreationdate, maxcreationdate, minfirstalbum, maxfirstalbum, location, members)
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	err = templ.Execute(w, &artist)
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}
