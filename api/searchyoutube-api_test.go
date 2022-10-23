package api

import (
	"main/structs"
	"testing"
)

var sampleSongInformation = structs.SongInformation{
	Id:         0,
	Title:      "Synchrograzer",
	Artist:     "Nana Mizuki",
	AnimeTitle: "Symphogear",
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
		name string
		args args
		want string
	}{
		{name: "checkTest", args: args{theme: sampleSongInformation}, want: "https://www.youtube.com/watch?v=2DKCoLZAGvQ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getYoutubeSongUrl(tt.args.theme); got != tt.want {
				t.Errorf("getYoutubeSongUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
