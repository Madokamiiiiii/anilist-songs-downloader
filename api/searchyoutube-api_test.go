package api

import (
	"main/structs"
	"testing"
)

func Test_getYoutubeSongUrl(t *testing.T) {
	type args struct {
		theme struct{ structs.AnimeInformations }
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "checkTest", args: args{theme: struct{ structs.AnimeInformations }{structs.Animethemes: structs.Animethemes{
			Type:     "OP",
			Sequence: 1,
			Song: struct {
				ID      int    `json:"id"`
				Title   string `json:"title"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
			}(struct {
				ID      int
				Title   string
				Artists []struct {
					Name string
				}
			}{
				ID:    123,
				Title: "Vitalization",
				Artists: []struct {
					Name string
				}{{Name: "Nana Mizuki"}},
			}),
		}}}, want: "https://youtube.com/watch?v=7lrw8VeBldc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYoutubeSongUrl(tt.args.theme); got != tt.want {
				t.Errorf("getYoutubeSongUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
