package api

import (
	"fmt"
	"github.com/animenotifier/anilist"
	"reflect"
)

func GetAnimesFromAniListUser(username string) []int {
	userId, error := anilist.GetUser(username)

	if error != nil {
		fmt.Errorf(error.Error())
	}

	animeList, error := anilist.GetAnimeList(userId.ID)

	if error != nil {
		fmt.Errorf(error.Error())
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
			fmt.Printf("No MAL ID for AL ID %v", listItem.Anime.ID)
		} else {
			animeIds = append(animeIds, listItem.Anime.MALID)
		}
	}

	return animeIds
}
