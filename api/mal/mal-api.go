package mal

import (
	"github.com/darenliang/jikan-go"
	"log"
	"time"
)

// This will be a fallback in case Animethemes doesn't know the anime

func GetSongInformationFromMAL(malId int) (jikan.AnimeThemes, error) {
	themes, err := jikan.GetAnimeThemes(malId)
	time.Sleep(1 * time.Second)

	if err != nil {
		log.Println(err.Error())
		return jikan.AnimeThemes{}, err
	}

	return *themes, nil
}

func GetEnglishTitle(id int) string {
	animeById, err := jikan.GetAnimeById(id)
	time.Sleep(1 * time.Second)

	if err != nil {
		log.Printf("MAL %v not found.\n", id)
		return ""
	}

	return animeById.Data.TitleEnglish
}

func GetTitle(id int) (string, error) {
	animeById, err := jikan.GetAnimeById(id)
	time.Sleep(1 * time.Second)

	if err != nil {
		log.Printf("MAL %v not found.\n", id)
		return "", err
	}

	return animeById.Data.Title, nil
}
