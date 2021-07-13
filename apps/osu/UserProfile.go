package osu

import "time"

type UserProfile struct {
	AvatarUrl     string      `json:"avatar_url"`
	CountryCode   string      `json:"country_code"`
	DefaultGroup  string      `json:"default_group"`
	Id            int         `json:"id"`
	IsActive      bool        `json:"is_active"`
	IsBot         bool        `json:"is_bot"`
	IsDeleted     bool        `json:"is_deleted"`
	IsOnline      bool        `json:"is_online"`
	IsSupporter   bool        `json:"is_supporter"`
	LastVisit     time.Time   `json:"last_visit"`
	PmFriendsOnly bool        `json:"pm_friends_only"`
	ProfileColour interface{} `json:"profile_colour"`
	Username      string      `json:"username"`
	CoverUrl      string      `json:"cover_url"`
	Discord       interface{} `json:"discord"`
	HasSupported  bool        `json:"has_supported"`
	Interests     string      `json:"interests"`
	JoinDate      time.Time   `json:"join_date"`
	Kudosu        struct {
		Total     int `json:"total"`
		Available int `json:"available"`
	} `json:"kudosu"`
	Location     string      `json:"location"`
	MaxBlocks    int         `json:"max_blocks"`
	MaxFriends   int         `json:"max_friends"`
	Occupation   string      `json:"occupation"`
	Playmode     string      `json:"playmode"`
	Playstyle    []string    `json:"playstyle"`
	PostCount    int         `json:"post_count"`
	ProfileOrder []string    `json:"profile_order"`
	Title        interface{} `json:"title"`
	TitleUrl     interface{} `json:"title_url"`
	Twitter      interface{} `json:"twitter"`
	Website      interface{} `json:"website"`
	Country      struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"country"`
	Cover struct {
		CustomUrl string      `json:"custom_url"`
		Url       string      `json:"url"`
		Id        interface{} `json:"id"`
	} `json:"cover"`
	AccountHistory           []interface{} `json:"account_history"`
	ActiveTournamentBanner   interface{}   `json:"active_tournament_banner"`
	Badges                   []interface{} `json:"badges"`
	BeatmapPlaycountsCount   int           `json:"beatmap_playcounts_count"`
	CommentsCount            int           `json:"comments_count"`
	FavouriteBeatmapsetCount int           `json:"favourite_beatmapset_count"`
	FollowerCount            int           `json:"follower_count"`
	GraveyardBeatmapsetCount int           `json:"graveyard_beatmapset_count"`
	Groups                   []interface{} `json:"groups"`
	LovedBeatmapsetCount     int           `json:"loved_beatmapset_count"`
	MappingFollowerCount     int           `json:"mapping_follower_count"`
	MonthlyPlaycounts        []struct {
		StartDate string `json:"start_date"`
		Count     int    `json:"count"`
	} `json:"monthly_playcounts"`
	Page struct {
		Html string `json:"html"`
		Raw  string `json:"raw"`
	} `json:"page"`
	PendingBeatmapsetCount int      `json:"pending_beatmapset_count"`
	PreviousUsernames      []string `json:"previous_usernames"`
	RankedBeatmapsetCount  int      `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts   []struct {
		StartDate string `json:"start_date"`
		Count     int    `json:"count"`
	} `json:"replays_watched_counts"`
	ScoresBestCount   int `json:"scores_best_count"`
	ScoresFirstCount  int `json:"scores_first_count"`
	ScoresRecentCount int `json:"scores_recent_count"`
	Statistics        struct {
		Level struct {
			Current  int `json:"current"`
			Progress int `json:"progress"`
		} `json:"level"`
		GlobalRank             int     `json:"global_rank"`
		Pp                     float64 `json:"pp"`
		RankedScore            int64   `json:"ranked_score"`
		HitAccuracy            float64 `json:"hit_accuracy"`
		PlayCount              int     `json:"play_count"`
		PlayTime               int     `json:"play_time"`
		TotalScore             int64   `json:"total_score"`
		TotalHits              int     `json:"total_hits"`
		MaximumCombo           int     `json:"maximum_combo"`
		ReplaysWatchedByOthers int     `json:"replays_watched_by_others"`
		IsRanked               bool    `json:"is_ranked"`
		GradeCounts            struct {
			Ss  int `json:"ss"`
			Ssh int `json:"ssh"`
			S   int `json:"s"`
			Sh  int `json:"sh"`
			A   int `json:"a"`
		} `json:"grade_counts"`
		CountryRank int `json:"country_rank"`
		Rank        struct {
			Country int `json:"country"`
		} `json:"rank"`
	} `json:"statistics"`
	SupportLevel     int `json:"support_level"`
	UserAchievements []struct {
		AchievedAt    time.Time `json:"achieved_at"`
		AchievementId int       `json:"achievement_id"`
	} `json:"user_achievements"`
	RankHistory struct {
		Mode string `json:"mode"`
		Data []int  `json:"data"`
	} `json:"rankHistory"`
	RankHistory1 struct {
		Mode string `json:"mode"`
		Data []int  `json:"data"`
	} `json:"rank_history"`
	RankedAndApprovedBeatmapsetCount int `json:"ranked_and_approved_beatmapset_count"`
	UnrankedBeatmapsetCount          int `json:"unranked_beatmapset_count"`
}
