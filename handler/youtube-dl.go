package handler

import (
	"context"
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
	for _, information := range informations {
		downloadSong(information)
	}
}

func downloadSong(information structs.SongInformation) {
	goutubedl.Path = "yt-dlp"

	result, err := goutubedl.New(context.Background(), information.Url, goutubedl.Options{Type: goutubedl.TypeSingle})
	if err != nil {
		log.Fatal(err)
	}

	downloadResult, err := result.Download(context.Background(), "best")

	if err != nil {
		log.Fatal(err)
	}
	defer downloadResult.Close()

	fileName := strings.Replace(result.Info.Title, "/", " ", -1)
	filePath := videoSaveDirectory + fileName + ".mp4"
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(f, downloadResult)
	if err != nil {
		return
	}

	err = ffmpeg_go.Input(filePath).Output(audioSaveDirectory+fileName+".mp3", ffmpeg_go.KwArgs{"q:a": 0, "map": "0:a:0", "af": "silenceremove=stop_periods=-1:stop_duration=1:stop_threshold=-50dB"}).OverWriteOutput().Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

	f.Close()
	downloadResult.Close()
	err = os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}

	SaveDownloaded(information.Id, true)
}
