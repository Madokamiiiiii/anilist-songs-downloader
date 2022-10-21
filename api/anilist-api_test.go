package api

import (
	"golang.org/x/exp/slices"
	"testing"
)

func Test_getAnimesFromAniListUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"checkForWatchedAnime", args{username: "Youmukami"}, 15793},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAnimesFromAniListUser(tt.args.username); slices.Contains(got, tt.want) {
				t.Errorf("getAnimesFromAniListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
