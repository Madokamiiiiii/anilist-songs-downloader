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

const BASE_URL = "https://api.animethemes.moe/anime?"
const FILTER = "filter[has]=resources&include=animethemes.song&include=animethemes.song.artists&fields[anime]=name&fields[artist]=name&fields[animetheme]=type,sequence&fields[song]=id,title&filter[site]=MyAnimeList&filter[external_id]=%v"

func GetSongsForIdList(anilistIds []int) []structs.AnimeInformations {
	var animeInformations []structs.AnimeInformations

	for i, id := range anilistIds {
		animeInformation, err := getSongsForAnime(id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			animeInformations = append(animeInformations, animeInformation)
		}

		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	return animeInformations
}

func getSongsForAnime(anilistId int) (structs.AnimeInformations, error) {
	requestURL := fmt.Sprintf(BASE_URL+FILTER, anilistId)

	response, err := http.Get(requestURL)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	information := structs.AnimeInformations{}

	err = json.Unmarshal(responseData, &information)
	if err != nil {
		fmt.Println("Could not decode response")
		os.Exit(2)
	}

	if len(information.Anime) == 0 {
		return information, fmt.Errorf("anime not found on Animethemes for the id %v", anilistId)
	}

	return information, nil
}
