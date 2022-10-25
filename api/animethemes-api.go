package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/structs"
	"net/http"
	"os"
)

const BaseUrl = "https://api.animethemes.moe/anime?"
const FILTER = "filter[has]=resources&include=animethemes.song&include=animethemes.song.artists&fields[anime]=name&fields[artist]=name&fields[animetheme]=type,sequence&fields[song]=id,title&filter[site]=MyAnimeList&filter[external_id]=%v"

func GetSongs(anilistInformation []structs.AniListInformation) ([]structs.AnimeInformations, []int) {
	var animeInformations []structs.AnimeInformations
	var malList []int

	for i, information := range anilistInformation {
		animeInformation, err := GetSongsForAnime(information)
		if err != nil {
			log.Println(err.Error())
			malList = append(malList, information.Id)
		} else {
			animeInformations = append(animeInformations, animeInformation)
		}

		log.Println(i)
	}

	return animeInformations, malList
}

func GetSongsForAnime(anilistInformation structs.AniListInformation) (structs.AnimeInformations, error) {
	requestURL := fmt.Sprintf(BaseUrl+FILTER, anilistInformation.Id)

	response, err := http.Get(requestURL)

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	information := structs.AnimeInformations{}

	err = json.Unmarshal(responseData, &information)
	if err != nil {
		log.Fatalln("Could not decode response")
	}

	if len(information.Anime) == 0 {
		return information, fmt.Errorf("anime not found on Animethemes for the id %v", anilistInformation.Id)
	}

	information.Anime[0].Id = anilistInformation.Id
	information.Anime[0].Name = anilistInformation.AnimeTitle

	return information, nil
}
