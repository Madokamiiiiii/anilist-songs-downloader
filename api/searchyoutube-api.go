package api

import (
	"fmt"
	"github.com/Pauloo27/searchtube"
	"main/structs"
	"os"
)

func GetYoutubeSongUrl(theme structs.SongInformation) string {

	title := theme.Title
	artist := theme.Artist
	searchTerm := title + " " + artist

	fmt.Println(searchTerm)

	searchResult, error := searchtube.Search(searchTerm, 1)
	if error != nil {
		error.Error()
		os.Exit(2)
	}

	return searchResult[0].URL
}
