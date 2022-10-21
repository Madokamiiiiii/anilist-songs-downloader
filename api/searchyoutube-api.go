package api

import (
	"fmt"
	"github.com/Pauloo27/searchtube"
	"main/handler"
	"main/structs"
	"os"
)

func getYoutubeSongUrl(theme structs.SongInformation) string {

	title := theme.Title
	artist := theme.Artist
	searchTerm := title + " " + artist

	fmt.Println(searchTerm)

	searchResult, error := searchtube.Search(searchTerm, 1)
	if error != nil {
		error.Error()
		os.Exit(2)
	}

	handler.SaveUrl(theme.Id, searchResult[0].URL)

	return searchResult[0].URL
}

func GetYoutubeSongUrlFromList(themes []structs.SongInformation) {
	for _, theme := range themes {
		getYoutubeSongUrl(theme)
	}
}
