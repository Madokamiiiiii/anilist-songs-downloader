package handler

import (
	"database/sql"
	"fmt"
	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/api/mal"
	"main/structs"
	"strconv"
	"strings"
)

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

		addToALDB(songInformation)
	}
}

func ConvertAnimethemesToSongInformation(informations []structs.AnimeInformations, malList []int) {

	for _, information := range informations {
		convertAnimethemeToSongInformation(information)
	}

	for _, id := range malList {
		getAndParseMALThemes(id)
	}
}

func getAndParseMALThemes(id int) {
	var information structs.SongInformation

	themes, err := mal.GetSongInformationFromMAL(id)
	if err != nil {
		themes, err = mal.GetSongInformationFromMAL(id)
		if err != nil {
			log.Printf("Skipping %v\n", id)
			return
		}
	}

	animeTitle, err := mal.GetTitle(id)
	if err != nil {
		animeTitle, err = mal.GetTitle(id)
		if err != nil {
			log.Printf("Skipping %v\n", id)
			return
		}
	}

	for i, opening := range themes.Data.Openings {
		splitted := strings.Split(opening, "by")
		title := strings.Replace(splitted[0], "\"", "", -1)
		title = strings.TrimLeft(title, "1234567890:")
		artist := ""
		if len(splitted) == 2 {
			artist = splitted[1]
			artist = strings.TrimRight(artist, "(eps-1234567890)")
		}
		information = structs.SongInformation{
			Id:         0,
			Title:      title,
			Artist:     artist,
			AnimeTitle: animeTitle,
			AnimeId:    strconv.Itoa(id),
			Type:       "OP",
			Sequence:   i + 1,
			Downloaded: false,
			Url:        "",
		}

		addToMALDB(information)
	}

	for i, ending := range themes.Data.Endings {
		splitted := strings.Split(ending, " by ")
		title := strings.Replace(splitted[0], "\"", "", -1)
		title = strings.TrimLeft(title, "1234567890:")
		artist := ""
		if len(splitted) == 2 {
			artist = splitted[1]
			artist = strings.TrimRight(artist, "(eps-1234567890)")
		}

		information = structs.SongInformation{
			Id:         0,
			Title:      title,
			Artist:     artist,
			AnimeTitle: animeTitle,
			AnimeId:    strconv.Itoa(id),
			Type:       "ED",
			Sequence:   i + 1,
			Downloaded: false,
			Url:        "",
		}

		addToMALDB(information)
	}
}

func InitDB() {
	database, _ := sql.Open("sqlite3", "./songs.db")
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS SongInformation (`id` INTEGER PRIMARY KEY, `title` VARCHAR(255) NOT NULL, `animetitle` VARCHAR(255) NOT NULL, `animeid` INTEGER, `artist` VARCHAR(255) NULL, `type` VARCHAR(20) NOT NULL, `sequence` INTEGER NOT NULL, `downloaded` SMALLINT, `url` VARCHAR)")
	checkErr(err)

	_, err = database.Exec("CREATE TABLE IF NOT EXISTS mallist (`title` VARCHAR(255) PRIMARY KEY, `animetitle` VARCHAR(255) NOT NULL, `animeid` INTEGER, `artist` VARCHAR(255) NULL, `type` VARCHAR(20) NOT NULL, `sequence` INTEGER NOT NULL, `downloaded` SMALLINT, `url` VARCHAR)")
	checkErr(err)

	db = database
}

func addToMALDB(information structs.SongInformation) {
	var exists bool
	if _ = db.QueryRow("SELECT EXISTS(SELECT 1 FROM mallist WHERE title=?)", information.Title).Scan(&exists); exists {
		log.Println(fmt.Sprintf("Entry %v already exists", information.Title))
	} else {
		_, err := db.Exec("INSERT INTO mallist(title, animetitle, animeid, artist, type, sequence, downloaded, url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", information.Title, information.AnimeTitle, information.AnimeId, information.Artist, information.Type, information.Sequence, information.Downloaded, information.Url)
		checkErr(err)
	}
}

func addToALDB(information structs.SongInformation) {
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
	var songInformationsAL []structs.SongInformation
	var songInformationsMAL []structs.SongInformation

	rows, err := db.Query("SELECT * FROM SongInformation")
	checkErr(err)

	err = scan.Rows(&songInformationsAL, rows)
	checkErr(err)

	rows, err = db.Query("SELECT * FROM mallist")
	checkErr(err)

	err = scan.Rows(&songInformationsMAL, rows)
	checkErr(err)

	return append(songInformationsMAL, songInformationsAL...)
}

func GetNotLoadedSongInformations() []structs.SongInformation {
	var songInformationsAL []structs.SongInformation
	var songInformationsMAL []structs.SongInformation

	rows, err := db.Query("SELECT * FROM SongInformation WHERE animetitle = '' OR url = ''")
	checkErr(err)

	err = scan.Rows(&songInformationsAL, rows)
	checkErr(err)

	rows, err = db.Query("SELECT * FROM mallist WHERE animetitle = '' OR url = ''")
	checkErr(err)

	err = scan.Rows(&songInformationsMAL, rows)
	checkErr(err)

	return append(songInformationsMAL, songInformationsAL...)
}

func SaveUrl(id int, url string) {
	var oldUrl string
	row, err := db.Query("SELECT url FROM SongInformation WHERE ID = ?", id)
	checkErr(err)

	err = scan.Row(&oldUrl, row)
	checkErr(err)

	if oldUrl == "" {

	} else if oldUrl != url {
		var confirm string

		log.Printf("New URL detected. Old: %v, New: %v\n", oldUrl, url)
		log.Println("Would you like to use the new video? [Y]/N")
		//	_, _ = fmt.Scanln(&confirm)

		if confirm == "n" || confirm == "N" {
			return
		}
		SaveDownloaded(id, false)
	} else if oldUrl == url {
		return
	}

	_, err = db.Exec("UPDATE SongInformation SET url = ? WHERE id = ?", url, id)
	checkErr(err)
}

func SaveMALUrl(title string, url string) {
	var oldUrl string
	row, err := db.Query("SELECT url FROM mallist WHERE title = ?", title)
	checkErr(err)

	err = scan.Row(&oldUrl, row)
	checkErr(err)

	if oldUrl == "" {

	} else if oldUrl != url {
		var confirm string

		log.Printf("New URL detected. Old: %v, New: %v\n", oldUrl, url)
		log.Println("Would you like to use the new video? [Y]/N")
		//	_, _ = fmt.Scanln(&confirm)

		if confirm == "n" || confirm == "N" {
			return
		}
		SaveMALDownloaded(title, false)
	} else if oldUrl == url {
		return
	}

	_, err = db.Exec("UPDATE mallist SET url = ? WHERE title = ?", url, title)
	checkErr(err)
}

func SaveDownloaded(id int, saved bool) {
	_, err := db.Exec("UPDATE SongInformation SET downloaded = ? WHERE id = ?", saved, id)
	checkErr(err)
}

func SaveMALDownloaded(title string, saved bool) {
	_, err := db.Exec("UPDATE mallist SET downloaded = ? WHERE title = ?", saved, title)
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

func Delete() {
	db.Exec("DELETE FROM SongInformation WHERE animetitle = '' OR url = ''")

}
