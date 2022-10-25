package api

import (
	"testing"
)

func Test_getSongsForAnime(t *testing.T) {
	type args struct {
		anilistId int
	}

	tests := []struct {
		name   string
		args   args
		wanted string
	}{
		{"checkURL", args{anilistId: 15793}, "Senki Zesshou Symphogear G: In the Distance, That Day, When the Star Became Music..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := GetSongsForAnime(tt.args.anilistId)
			if result.Anime[0].Name != tt.wanted {
				t.Errorf("wanted: %v, got: %v", tt.wanted, result.Anime[0].Name)
			}
		})
	}
}

func Test_getSongsForIdList(t *testing.T) {
	type args struct {
		anilistIds []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"TestGetSongs", args{anilistIds: []int{15793, 5420}}, []string{"Senki Zesshou Symphogear G: In the Distance, That Day, When the Star Became Music...", "Kemono no Souja Erin"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSongsForIdList(tt.args.anilistIds)
			if len(got) != 2 {
				t.Errorf("getSongsForIdList() = %v, want %v", got, tt.want)
			} else {
				for i := range got {
					if got[i].Anime[0].Name != tt.want[i] {
						t.Errorf("getSongsForIdList() = %v, want %v", got[i].Anime[0].Name, tt.want[i])
					}
				}
			}
		})
	}
}
