package handler

import (
	"database/sql"
	"fmt"
	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/structs"
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
			Id:         animetheme.Song.ID,
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

func ConvertAnimethemesToSongInformation(informations []structs.AnimeInformations) {

	for _, information := range informations {
		convertAnimethemeToSongInformation(information)
	}
}

func AddToMALList(id int) {
	malList = append(malList, id)
}

func InitDB() {
	database, _ := sql.Open("sqlite3", "./songs.db")
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS SongInformation (`id` INTEGER PRIMARY KEY, `title` VARCHAR(255) NOT NULL, `animetitle` VARCHAR(255) NOT NULL, `artist` VARCHAR(255) NULL, `type` VARCHAR(20) NOT NULL, `sequence` INTEGER NOT NULL, `downloaded` SMALLINT, `url` VARCHAR)")
	checkErr(err)

	db = database
}

func addToDB(information structs.SongInformation) {
	var exists bool
	if _ = db.QueryRow("SELECT EXISTS(SELECT 1 FROM SongInformation WHERE id=?)", information.Id).Scan(&exists); exists {
		log.Println(fmt.Sprintf("Entry %v already exists", information.Id))
	} else {
		_, err := db.Exec("INSERT INTO SongInformation(id, title, animetitle, artist, type, sequence, downloaded, url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", information.Id, information.Title, information.AnimeTitle, information.Artist, information.Type, information.Sequence, information.Downloaded, information.Url)
		checkErr(err)
	}
}

func getSongInformation(ID int) structs.SongInformation {
	var songInformation structs.SongInformation

	row, err := db.Query("SELECT * FROM SongInformation WHERE Id=?", ID)
	checkErr(err)

	err = scan.Row(&songInformation, row)
	checkErr(err)

	return songInformation
}

func GetAllSongInformations() []structs.SongInformation {
	var songInformations []structs.SongInformation

	rows, err := db.Query("SELECT * FROM SongInformation")
	checkErr(err)

	err = scan.Rows(&songInformations, rows)
	checkErr(err)

	return songInformations
}

func SaveUrl(id int, url string) {
	_, err := db.Exec("UPDATE SongInformation SET url = ? WHERE id = ?", url, id)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
