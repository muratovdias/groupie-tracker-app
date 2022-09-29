package utils

import (
	"strconv"
	"strings"

	"tracker/internal/model"
)

func SearchGroup(data string) (*model.AllArtist, error) {
	var result []model.Artist
	artists, err := GetData()
	if err != nil {
		return &model.AllArtist{}, err
	}
	for _, artist := range artists.Artists {
		var containe bool
		if myContains(artist.Name, data) {
			result = append(result, artist)
			continue
		} else if strings.Contains(strconv.Itoa(artist.CreationDate), data) {
			result = append(result, artist)
			continue
		} else if strings.Contains(artist.FirstAlbum, data) {
			result = append(result, artist)
			continue
		} else {
			for _, m := range artist.Members {
				if strings.Contains(m, data) {
					result = append(result, artist)
					containe = true
					break
				}
			}
			if containe {
				continue
			}
		}

		for city, dates := range artist.DatesLocations {
			if myContains(city, data) {
				result = append(result, artist)
				break
			}
			for _, date := range dates {
				if myContains(date, data) {
					result = append(result, artist)
					containe = true
					break
				}
			}
			if containe {
				break
			}
		}
	}
	allArtists := model.AllArtist{Artists: artists.Artists, FoundArtists: result, SearchText: data}

	return &allArtists, nil
}

// function for search in anycase
func myContains(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
