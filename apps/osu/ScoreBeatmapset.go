package osu

type ScoreBeatmapset struct {
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers        struct {
		Cover       string `json:"cover"`
		Cover2X     string `json:"cover@2x"`
		Card        string `json:"card"`
		Card2X      string `json:"card@2x"`
		List        string `json:"list"`
		List2X      string `json:"list@2x"`
		Slimcover   string `json:"slimcover"`
		Slimcover2X string `json:"slimcover@2x"`
	} `json:"covers"`
	Creator        string `json:"creator"`
	FavouriteCount int    `json:"favourite_count"`
	Hype           struct {
		Current  int `json:"current"`
		Required int `json:"required"`
	} `json:"hype"`
	Id           int    `json:"id"`
	Nsfw         bool   `json:"nsfw"`
	PlayCount    int    `json:"play_count"`
	PreviewUrl   string `json:"preview_url"`
	Source       string `json:"source"`
	Status       string `json:"status"`
	Title        string `json:"title"`
	TitleUnicode string `json:"title_unicode"`
	UserId       int    `json:"user_id"`
	Video        bool   `json:"video"`
}
