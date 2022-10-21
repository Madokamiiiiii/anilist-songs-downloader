package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/exp/slices"
	"main/structs"
	"os"
)

var songInformations []structs.SongInformation

var malList []int

var db *sql.DB

func convertAnimethemeToSongInformation(informations structs.AnimeInformations) {

	animeTitle := informations.Anime[0].Name

	for _, animetheme := range informations.Anime[0].Animethemes {
		var artist string
		if len(animetheme.Song.Artists) == 0 {
			artist = ""
		} else {
			artist = animetheme.Song.Artists[0].Name
		}

		songInformation := structs.SongInformation{
			ID:         animetheme.Song.ID,
			Title:      animetheme.Song.Title,
			Artist:     artist,
			AnimeTitle: animeTitle,
			Type:       animetheme.Type,
			Sequence:   animetheme.Sequence,
		}

		addToDB(songInformation)

		songInformations = append(songInformations, songInformation)
	}
}

func ConvertAndSaveAnimethemesToSongInformation(informations []structs.AnimeInformations) {
	initDB()

	for _, information := range informations {
		convertAnimethemeToSongInformation(information)
	}

	saveToFile()
}

func AddToMALList(id int) {
	malList = append(malList, id)
}

func saveToFile() {
	var savedInformations []structs.SongInformation

	savedFile, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err = json.Unmarshal(savedFile, &savedInformations)
	}

	var existingSongIds []int
	for _, information := range savedInformations {
		existingSongIds = append(existingSongIds, information.ID)
	}

	for _, information := range songInformations {
		if !slices.Contains(existingSongIds, information.ID) {
			savedInformations = append(savedInformations, information)
		}
	}

	file, _ := json.MarshalIndent(savedInformations, "", " ")

	err = os.WriteFile("test.json", file, 0644)
	checkErr(err)
}

func initDB() {
	db, err := sql.Open("sqlite3", "./songs.db")
	if err != nil {
		os.Create("./songs.db")
		db, err = sql.Open("sqlite3", "./songs.db")
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS SongInformation (`id` INTEGER PRIMARY KEY, `title` VARCHAR(255) NOT NULL, `artist` VARCHAR(255) NULL, `type` VARCHAR(20) NOT NULL, `sequence` INTEGER NOT NULL, `downloaded` SMALLINT, `url` VARCHAR)")
	}
}

func addToDB(information structs.SongInformation) {
	var songInformation structs.SongInformation
	err := db.QueryRow("SELECT ID FROM SongInformation WHERE ID=" + string(rune(information.ID))).Scan(&songInformation.ID)
	checkErr(err)

	if songInformation.ID == 0 {
		db.Exec("INSERT INTO SongInformation")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
