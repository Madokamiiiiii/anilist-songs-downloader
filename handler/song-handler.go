package handler

import (
	"database/sql"
	"fmt"
	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/structs"
	"strconv"
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
			AnimeId:    strconv.Itoa(informations.Anime[0].Id),
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
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS SongInformation (`id` INTEGER PRIMARY KEY, `title` VARCHAR(255) NOT NULL, `animetitle` VARCHAR(255) NOT NULL, `animeid` INTEGER, `artist` VARCHAR(255) NULL, `type` VARCHAR(20) NOT NULL, `sequence` INTEGER NOT NULL, `downloaded` SMALLINT, `url` VARCHAR)")
	checkErr(err)

	db = database
}

func addToDB(information structs.SongInformation) {
	var exists bool
	if _ = db.QueryRow("SELECT EXISTS(SELECT 1 FROM SongInformation WHERE id=?)", information.Id).Scan(&exists); exists {
		log.Println(fmt.Sprintf("Entry %v already exists", information.Id))
	} else {
		_, err := db.Exec("INSERT INTO SongInformation(id, title, animetitle, animeid, artist, type, sequence, downloaded, url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", information.Id, information.Title, information.AnimeTitle, information.AnimeId, information.Artist, information.Type, information.Sequence, information.Downloaded, information.Url)
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
	var oldUrl string
	row, err := db.Query("SELECT url FROM SongInformation WHERE ID = ?", id)
	checkErr(err)

	err = scan.Row(&oldUrl, row)
	checkErr(err)

	if oldUrl == url {
		return
	} else if oldUrl != url {
		var confirm string

		log.Printf("New URL detected. Old: %v, New: %v\n", oldUrl, url)
		log.Println("Would you like to use the new video? [Y]/N")
		_, err := fmt.Scanln(&confirm)
		if err != nil {
			checkErr(err)
		}

		if confirm == "n" || confirm == "N" {
			return
		}
		SaveDownloaded(id, false)
	}

	_, err = db.Exec("UPDATE SongInformation SET url = ? WHERE id = ?", url, id)
	checkErr(err)
}

func SaveDownloaded(id int, saved bool) {
	_, err := db.Exec("UPDATE SongInformation SET downloaded = ? WHERE id = ?", saved, id)
	checkErr(err)
}

func GetNotDownloadedSongs() []structs.SongInformation {
	var songInformations []structs.SongInformation

	rows, err := db.Query("SELECT * FROM SongInformation WHERE downloaded=0")
	checkErr(err)

	err = scan.Rows(&songInformations, rows)
	checkErr(err)

	return songInformations
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
