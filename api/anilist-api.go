package api

import (
	"fmt"
	"github.com/animenotifier/anilist"
	"log"
	"reflect"
)

func GetAnimesFromAniListUser(username string) []int {
	userId, err := anilist.GetUser(username)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	animeList, err := anilist.GetAnimeList(userId.ID)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	var completedList []*anilist.AnimeListItem
	fmt.Println(reflect.TypeOf(completedList))

	for _, list := range animeList.Lists {
		if list.Name == "Completed" {
			completedList = list.Entries
			break
		}
	}

	var animeIds []int

	for _, listItem := range completedList {
		if listItem.Anime.MALID == 0 {
			log.Printf("No MAL Id for AL Id %v\n", listItem.Anime.ID)
		} else {
			animeIds = append(animeIds, listItem.Anime.MALID)
		}
	}

	return animeIds
}
