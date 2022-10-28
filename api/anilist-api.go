package api

import (
	"fmt"
	"github.com/animenotifier/anilist"
	"log"
	"main/structs"
	"reflect"
)

func GetAnimesFromAniListUser(username string) []structs.AniListInformation {
	var anilistInformation []structs.AniListInformation
	userId, err := anilist.GetUser(username)

	if err != nil {
		err = fmt.Errorf(err.Error())
		log.Fatalln(err)
	}

	animeList, err := anilist.GetAnimeList(userId.ID)

	if err != nil {
		err = fmt.Errorf(err.Error())
		log.Fatalln(err)
	}

	var completedList []*anilist.AnimeListItem
	fmt.Println(reflect.TypeOf(completedList))

	for _, list := range animeList.Lists {
		if list.Name == "Completed" {
			completedList = list.Entries
			break
		}
	}

	for _, listItem := range completedList {
		if listItem.Anime.MALID == 0 {
			log.Printf("No MAL Id for AL Id %v\n", listItem.Anime.ID)
		} else {
			var title string
			if listItem.Anime.Title.English == "" {
				title = listItem.Anime.Title.Romaji
			} else {
				title = listItem.Anime.Title.English
			}
			anilistInformation = append(anilistInformation, structs.AniListInformation{
				Id:         listItem.Anime.MALID,
				AnimeTitle: title,
			})
		}
	}

	return anilistInformation
}
