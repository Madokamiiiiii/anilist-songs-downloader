package api

import (
	"fmt"
	"github.com/darenliang/jikan-go"
	"os"
)

// This will be a fallback in case Animethemes doesn't know the anime

func getSongInformationFromMAL(malId int) {
	themes, error := jikan.GetAnimeThemes(malId)
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(2)
	}
	fmt.Println(themes)
}

func getSongInformationFromMALIds(malIds []int) {
	for _, id := range malIds {
		getSongsForAnime(id)
	}
}
