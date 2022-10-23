package main

import (
	"log"
	"main/api"
	"main/handler"
)

func main() {
	handler.InitDB()

	log.Println("Getting animeIds from AniList...")
	aniListIds := api.GetAnimesFromAniListUser("Youmukami")

	log.Println("Getting song informations from Animethemes...")
	res := api.GetSongsForIdList(aniListIds)

	log.Println("Converting and saving...")
	handler.ConvertAnimethemesToSongInformation(res)

	log.Println("Getting urls...")
	api.GetYoutubeSongUrlFromList(handler.GetAllSongInformations())

	log.Println("Downloading...")
	handler.DownloadSongs(handler.GetNotDownloadedSongs())
}
