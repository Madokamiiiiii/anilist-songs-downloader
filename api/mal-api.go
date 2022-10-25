package api

import (
	"github.com/darenliang/jikan-go"
	"log"
	"os"
)

// This will be a fallback in case Animethemes doesn't know the anime

func getSongInformationFromMAL(malId int) {
	themes, err := jikan.GetAnimeThemes(malId)
	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}
	log.Println(themes)
}

func getSongInformationFromMALIds(malIds []int) {
	for _, id := range malIds {
		getSongsForAnime(id)
	}
}

func GetEnglishTitle(id int) string {
	animeById, err := jikan.GetAnimeById(id)
	if err != nil {
		log.Printf("MAL %v not found.\n", id)
		return ""
	}

	return animeById.Data.TitleEnglish
}
