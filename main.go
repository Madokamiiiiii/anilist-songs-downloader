package main

import (
	"fmt"
	"main/api"
	"main/handler"
)

func main() {

	fmt.Println("Getting animeIds from AniList...")
	aniListIds := api.GetAnimesFromAniListUser("Youmukami")

	fmt.Println("Getting song informations from Animethemes...")
	res := api.GetSongsForIdList(aniListIds)

	fmt.Println("Converting and saving...")
	handler.ConvertAndSaveAnimethemesToSongInformation(res)

	fmt.Println("Downloading...")
	// downloadSongsFromAnime(res[0])
}
