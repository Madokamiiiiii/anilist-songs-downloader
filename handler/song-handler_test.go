package handler

import (
	"main/structs"
	"reflect"
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

var sampleSongInformation2 = structs.SongInformation{
	Id:         1,
	Title:      "Vitalization",
	Artist:     "Nana Mizuki",
	AnimeTitle: "Symphogear G",
	Type:       "OP",
	Sequence:   1,
	Downloaded: false,
	Url:        "",
}

func Test_initDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDB()
		})
	}
}

func Test_addToDB(t *testing.T) {
	type args struct {
		information structs.SongInformation
	}

	InitDB()

	tests := []struct {
		name string
		args args
	}{
		{"Test add song", args{information: sampleSongInformation}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addToDB(tt.args.information)
		})
	}
}

func Test_getSongInformation(t *testing.T) {
	InitDB()
	addToDB(sampleSongInformation)

	t.Run("TestCorrectSave", func(t *testing.T) {
		if got := getSongInformation(0); !reflect.DeepEqual(got, sampleSongInformation) {
			t.Errorf("getSongInformation() = %v, want %v", got, sampleSongInformation)
		}
	})
}

func Test_getAllSongInformations(t *testing.T) {
	InitDB()
	addToDB(sampleSongInformation)
	addToDB(sampleSongInformation2)

	var wantedInformations = []structs.SongInformation{sampleSongInformation, sampleSongInformation2}

	t.Run("TestCorrectSave", func(t *testing.T) {
		if got := GetAllSongInformations(); !reflect.DeepEqual(got, wantedInformations) {
			t.Errorf("getSongInformation() = %v, want %v", got, wantedInformations)
		}
	})
}

func TestSaveUrl(t *testing.T) {
	url := "test.com"

	InitDB()
	addToDB(sampleSongInformation)
	SaveUrl(0, url)

	if got := getSongInformation(0); got.Url != url {
		t.Errorf("getSongInformation() = %v, want %v", got.Url, url)
	}

}
