package api

import (
	"fmt"
	"github.com/Pauloo27/searchtube"
	"main/handler"
	"main/structs"
	"strconv"
	"testing"
)

var sampleSongInformation = structs.SongInformation{
	Id:         0,
	Title:      "Vitalization",
	Artist:     "Nana Mizuki",
	AnimeTitle: "Symphogear G",
	Type:       "OP",
	Sequence:   1,
	Downloaded: false,
	Url:        "",
}

var sampleSongInformationWithoutArtist = structs.SongInformation{
	Id:         1,
	Title:      "Connect",
	Artist:     "",
	AnimeTitle: "Madoka",
	Type:       "OP",
	Sequence:   1,
	Downloaded: false,
	Url:        "",
}

var sampleSongInformationThirdOption = structs.SongInformation{
	Id:         2,
	Title:      "10 hour Microwave sounds",
	Artist:     "",
	AnimeTitle: "",
	Type:       "OP",
	Sequence:   1,
	Downloaded: false,
	Url:        "",
}

func Test_getYoutubeSongUrl(t *testing.T) {
	type args struct {
		theme structs.SongInformation
	}
	tests := []struct {
		name         string
		args         args
		wantedSearch string
	}{
		{name: "checkTestFirstOption", args: args{theme: sampleSongInformation}, wantedSearch: sampleSongInformation.Title + " " + sampleSongInformation.Artist},
		{name: "checkTestSecondOption", args: args{theme: sampleSongInformationWithoutArtist}, wantedSearch: sampleSongInformationWithoutArtist.Title + " " + sampleSongInformationWithoutArtist.AnimeTitle},
		{name: "checkTestThirdOption", args: args{theme: sampleSongInformationThirdOption}, wantedSearch: sampleSongInformationThirdOption.AnimeId + " " + sampleSongInformationThirdOption.Type + " " + strconv.Itoa(sampleSongInformationThirdOption.Sequence) + " Full"},
	}

	handler.InitDB()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wanted, _ := searchtube.Search(tt.wantedSearch, 1)
			fmt.Println(tt.wantedSearch)
			if got := getYoutubeSongUrl(tt.args.theme); got != wanted[0].URL {
				t.Errorf("getYoutubeSongUrl() = %v, want %v", got, wanted[0].URL)
			}
		})
	}
}
