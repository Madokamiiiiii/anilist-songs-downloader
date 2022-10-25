package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"main/api"
	"main/handler"
	"os"
)

func main() {
	handler.InitDB()

	app := &cli.App{
		Name:  "anilist-songs-downloader",
		Usage: "Downloads full theme songs from an anilist from animethemes and youtube!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Value:    "Youmukami",
				Usage:    "anilist username",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "rescan",
				Value: false,
				Usage: "whether to fetch information again from animethemes",
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Hello friend!")

			log.Println("Getting animeIds from AniList...")
			aniListIds := api.GetAnimesFromAniListUser(cCtx.String("username"))

			log.Println("Getting song informations from Animethemes...")
			res, malList := api.GetSongs(aniListIds)

			log.Println("Converting and saving...")
			handler.ConvertAnimethemesToSongInformation(res, malList)

			log.Println("Getting urls...")
			api.GetYoutubeSongUrlFromList(handler.GetAllSongInformations())

			log.Println("Downloading...")
			handler.DownloadSongs(handler.GetNotDownloadedSongs())

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
