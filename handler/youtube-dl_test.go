package handler

import (
	"main/structs"
	"testing"
)

func TestDownloadSong(t *testing.T) {
	type args struct {
		informations structs.SongInformation
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testDownload", args: args{informations: sampleSongInformation}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadSong(tt.args.informations)
		})
	}
}
