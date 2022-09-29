package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tracker/internal/model"
)

var (
	artists          model.AllArtist
	relation         model.Relations
	jsonlinkArtists  = "https://groupietrackers.herokuapp.com/api/artists"
	jsonlinkRelation = "https://groupietrackers.herokuapp.com/api/relation"
)

func getJson(url string, object interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dates, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	jsonErr := json.Unmarshal(dates, &object)
	if jsonErr != nil {
		return err
	}
	return nil
}

func GetData() (*model.AllArtist, error) {
	if err := getJson(jsonlinkArtists, &artists.Artists); err != nil {
		return &artists, err
	}
	if err := getJson(jsonlinkRelation, &relation); err != nil {
		return &artists, err
	}
	for i, w := range relation.Location {
		artists.Artists[i].DatesLocations = w.DatesLocations
	}
	return &artists, nil
}
