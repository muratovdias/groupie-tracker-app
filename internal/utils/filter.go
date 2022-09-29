package utils

import (
	"strconv"
	"strings"
	"tracker/internal/model"
)

func Filter_artists(creationDate_min, creationDate_max string, album_min, album_max string, location string, members []string) (*model.AllArtist, error) {
	var artist []model.Artist
	creationDate_maxs, _ := strconv.Atoi(creationDate_max)
	creationDate_mins, _ := strconv.Atoi(creationDate_min)
	album_maxs, _ := strconv.Atoi(album_max)
	album_mins, _ := strconv.Atoi(album_min)
	artists, _ := GetData()
	for _, art := range artists.Artists {
		album, _ := strconv.Atoi(art.FirstAlbum[6:])
		var flag1, flag2, flag3 bool
		if (art.CreationDate <= creationDate_maxs && art.CreationDate >= creationDate_mins) && (album <= album_maxs && album >= album_mins) {
			flag1 = true
		}
		if len(location) == 0 {
			flag2 = true
		}

		for city := range art.DatesLocations {
			city = strings.Replace(strings.Replace(city, "_", " ", -1), "-", ", ", -1)
			if myContains(city, location) {
				flag2 = true
			}
		}
		for _, w := range members {
			s, _ := strconv.Atoi(w)
			if len(art.Members) == s {
				flag3 = true
			}
		}
		if len(members) == 0 {
			flag3 = true
		}
		if flag1 && flag2 && flag3 {
			artist = append(artist, art)
		}
	}
	allArtist := model.AllArtist{Artists: artists.Artists, FoundArtists: artist}
	return &allArtist, nil
}
