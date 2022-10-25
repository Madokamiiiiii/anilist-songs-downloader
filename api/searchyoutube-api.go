package api

import (
	"github.com/Pauloo27/searchtube"
	"log"
	"main/handler"
	"main/structs"
	"strconv"
	"time"
)

func getYoutubeSongUrl(theme structs.SongInformation) string {
	var searchResult []*searchtube.SearchResult
	var duration time.Duration = 0
	title := theme.Title
	// var animeId, _ = strconv.ParseInt(theme.AnimeId, 10, 0)
	// animeTitle := mal.GetEnglishTitle(int(animeId))
	animeTitle := theme.AnimeTitle
	// if animeTitle == "" {
	//	animeTitle = theme.AnimeTitle
	// }

	artist := theme.Artist

	if artist != "" {
		searchTerm := title + " " + artist
		searchResult, _ = searchtube.Search(searchTerm, 1)
		log.Println("Opt 1: " + title + " " + artist + " " + searchResult[0].URL)
		duration, _ = searchResult[0].GetDuration()
	}

	if duration.Minutes() > 10 || duration.Seconds() < 120 {

		searchTerm := title + " " + animeTitle
		searchResult, _ = searchtube.Search(searchTerm, 1)
		log.Println("Opt 2: " + searchTerm + " " + searchResult[0].URL)
		duration, _ = searchResult[0].GetDuration()

		if duration.Minutes() > 10 || duration.Seconds() < 120 {
			songType := theme.Type
			sequence := theme.Sequence
			if sequence != 0 {
				searchTerm = animeTitle + " " + songType + " " + strconv.Itoa(sequence) + " Full"
			} else {
				searchTerm = animeTitle + " " + songType + " Full"
			}

			searchResult, _ = searchtube.Search(searchTerm, 1)
			log.Println("Opt 3: " + searchTerm + " " + searchResult[0].URL)
		}

	}
	if theme.Id == 0 {
		handler.SaveMALUrl(theme.Title, searchResult[0].URL)
	} else {
		handler.SaveUrl(theme.Id, searchResult[0].URL)
	}

	return searchResult[0].URL
}

func GetYoutubeSongUrlFromList(themes []structs.SongInformation) {
	for _, theme := range themes {
		getYoutubeSongUrl(theme)
	}
}
