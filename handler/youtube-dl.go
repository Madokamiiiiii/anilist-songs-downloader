package handler

import "main/structs"

const saveDirectory = "songs"

func DownloadSongs(informations []structs.SongInformation) {
	for _, information := range informations {
		downloadSong(information)
	}
}

func downloadSong(information structs.SongInformation) {

}
