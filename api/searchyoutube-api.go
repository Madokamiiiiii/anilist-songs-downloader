package api

import (
	"fmt"
	"github.com/Pauloo27/searchtube"
	"main/handler"
	"main/structs"
	"os"
	"time"
)

func getYoutubeSongUrl(theme structs.SongInformation) string {

	title := theme.Title
	// artist := theme.Artist
	// searchTerm := title + " " + artist

	fmt.Println(title) // Just the title seems to be more accurate

	searchResult, err := searchtube.Search(title, 1)
	if err != nil {
		err.Error()
		os.Exit(2)
	}

	if searchResult[0].GetDuration() > time.ParseDuration("10m") {

	}

	handler.SaveUrl(theme.Id, searchResult[0].URL)

	return searchResult[0].URL
}

func GetYoutubeSongUrlFromList(themes []structs.SongInformation) {
	for _, theme := range themes {
		getYoutubeSongUrl(theme)
	}
}
