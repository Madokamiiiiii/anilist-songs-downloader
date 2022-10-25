package structs

type Animethemes struct {
	Type     string `json:"type"`
	Sequence int    `json:"sequence"`
	Song     struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
	} `json:"song"`
}

type AnimeInformations struct {
	Anime []struct {
		Name        string `json:"name"`
		Id          int
		Animethemes []Animethemes `json:"animethemes"`
	} `json:"anime"`
	Links struct {
		First string      `json:"first"`
		Last  interface{} `json:"last"`
		Prev  interface{} `json:"prev"`
		Next  interface{} `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
	} `json:"meta"`
}

type SongInformation struct {
	Id         int
	Title      string
	Artist     string
	AnimeTitle string `db:"animetitle"`
	AnimeId    string `db:"animeid"`
	Type       string
	Sequence   int
	Downloaded bool
	Url        string
}

type AniListInformation struct {
	Id         int
	AnimeTitle string
}
