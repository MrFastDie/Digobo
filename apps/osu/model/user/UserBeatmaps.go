package user

import "time"

type UserBeatmaps struct {
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
	Availability struct {
		DownloadDisabled bool        `json:"download_disabled"`
		MoreInformation  interface{} `json:"more_information"`
	} `json:"availability"`
	Bpm                float64   `json:"bpm"`
	CanBeHyped         bool      `json:"can_be_hyped"`
	DiscussionEnabled  bool      `json:"discussion_enabled"`
	DiscussionLocked   bool      `json:"discussion_locked"`
	IsScoreable        bool      `json:"is_scoreable"`
	LastUpdated        time.Time `json:"last_updated"`
	LegacyThreadUrl    string    `json:"legacy_thread_url"`
	NominationsSummary struct {
		Current  int `json:"current"`
		Required int `json:"required"`
	} `json:"nominations_summary"`
	Ranked        int         `json:"ranked"`
	RankedDate    interface{} `json:"ranked_date"`
	Storyboard    bool        `json:"storyboard"`
	SubmittedDate time.Time   `json:"submitted_date"`
	Tags          string      `json:"tags"`
	Beatmaps      []struct {
		BeatmapsetId     int         `json:"beatmapset_id"`
		DifficultyRating float64     `json:"difficulty_rating"`
		Id               int         `json:"id"`
		Mode             string      `json:"mode"`
		Status           string      `json:"status"`
		TotalLength      int         `json:"total_length"`
		UserId           int         `json:"user_id"`
		Version          string      `json:"version"`
		Accuracy         float64     `json:"accuracy"`
		Ar               float64     `json:"ar"`
		Bpm              float64     `json:"bpm"`
		Convert          bool        `json:"convert"`
		CountCircles     int         `json:"count_circles"`
		CountSliders     int         `json:"count_sliders"`
		CountSpinners    int         `json:"count_spinners"`
		Cs               float64     `json:"cs"`
		DeletedAt        interface{} `json:"deleted_at"`
		Drain            float64     `json:"drain"`
		HitLength        int         `json:"hit_length"`
		IsScoreable      bool        `json:"is_scoreable"`
		LastUpdated      time.Time   `json:"last_updated"`
		ModeInt          int         `json:"mode_int"`
		Passcount        int         `json:"passcount"`
		Playcount        int         `json:"playcount"`
		Ranked           int         `json:"ranked"`
		Url              string      `json:"url"`
		Checksum         string      `json:"checksum"`
	} `json:"beatmaps"`
}
