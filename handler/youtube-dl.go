package handler

import (
	"main/api"
	"main/structs"
)

func downloadSongsFromAnime(information structs.SongInformation) {

	url := api.GetYoutubeSongUrl(information)

	println(url)
}
