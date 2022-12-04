package handler

import (
	"context"
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"github.com/wader/goutubedl"
	"io"
	"log"
	"main/structs"
	"os"
	"strings"
)

const videoSaveDirectory = "./video/"
const audioSaveDirectory = "./audio/"

func DownloadSongs(informations []structs.SongInformation) {
	initFolders()
	for _, information := range informations {
		downloadSong(information)
	}
}

func initFolders() {
	_ = os.Mkdir(videoSaveDirectory, 0750)
	_ = os.Mkdir(audioSaveDirectory, 0750)
}

func downloadSong(information structs.SongInformation) {
	goutubedl.Path = "yt-dlp"

	result, err := goutubedl.New(context.Background(), information.Url, goutubedl.Options{Type: goutubedl.TypeSingle})
	if err != nil {
		log.Printf("Skipping: %v\n", information)
		return
	}

	downloadResult, err := result.Download(context.Background(), "best")

	if err != nil {
		log.Fatalln(err)
	}
	defer downloadResult.Close()

	var fileName string
	if information.Artist != "" {
		if information.Sequence == 0 {
			fileName = fmt.Sprintf("%v %v - %v by %v", information.AnimeTitle, information.Type, information.Title, information.Artist)
		} else {
			fileName = fmt.Sprintf("%v %v %v - %v by %v", information.AnimeTitle, information.Type, information.Sequence, information.Title, information.Artist)
		}
	} else {
		if information.Sequence == 0 {
			fileName = fmt.Sprintf("%v %v - %v", information.AnimeTitle, information.Type, information.Title)
		} else {
			fileName = fmt.Sprintf("%v %v %v - %v", information.AnimeTitle, information.Type, information.Sequence, information.Title)
		}
	}

	fileName = strings.Replace(fileName, "/", " ", -1)
	fileName = strings.Replace(fileName, "\\", " ", -1)
	fileName = strings.Replace(fileName, "*", " ", -1)
	fileName = strings.Replace(fileName, "|", "", -1)
	fileName = strings.Replace(fileName, "?", "", -1)
	fileName = strings.Replace(fileName, ":", "", -1)
	fileName = strings.Replace(fileName, "\"", "", -1)

	if len(fileName) > 100 {
		fileName = fileName[:100] // Too long for Windows
	}

	filePath := videoSaveDirectory + fileName + ".mp4"
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(f, downloadResult)
	if err != nil {
		return
	}

	err = ffmpeg_go.Input(filePath).Output(audioSaveDirectory+fileName+".mp3", ffmpeg_go.KwArgs{"q:a": 0, "map": "0:a:0", "af": "silenceremove=stop_periods=-1:stop_duration=1:stop_threshold=-65dB"}).OverWriteOutput().Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

	f.Close()
	downloadResult.Close()
	err = os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}

	if information.Id == 0 {
		SaveMALDownloaded(information.Title, true)
	} else {
		SaveDownloaded(information.Id, true)
	}
}
