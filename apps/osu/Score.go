package osu

import "time"

type ScoreMode string

const (
	SCORE_MODE_OSU    = "osu"    // osu!standard
	SCORE_MODE_FRUITS = "fruits" // osu!catch
	SCORE_MODE_MANIA  = "mania"  // osu!mania
	SCORE_MODE_TAIKO  = "taiko"  // osu!taiko
)

type ScoreType string

const (
	SCORE_TYPE_BEST   = "best"
	SCORE_TYPE_FIRSTS = "firsts"
	SCORE_TYPE_RECENT = "recent"
)

type Score struct {
	Id         int64           `json:"id"`
	UserId     int             `json:"user_id"`
	Accuracy   float64         `json:"accuracy"`
	Mods       []string        `json:"mods"`
	Score      int             `json:"score"`
	MaxCombo   int             `json:"max_combo"`
	Passed     bool            `json:"passed"`
	Perfect    bool            `json:"perfect"`
	Statistics ScoreStatistics `json:"statistics"`
	Rank       string          `json:"rank"`
	CreatedAt  time.Time       `json:"created_at"`
	BestId     interface{}     `json:"best_id"`
	Pp         interface{}     `json:"pp"`
	Mode       string          `json:"mode"`
	ModeInt    int             `json:"mode_int"`
	Replay     bool            `json:"replay"`
	Beatmap    ScoreBeatmap    `json:"beatmap"`
	Beatmapset ScoreBeatmapset `json:"beatmapset"`
	User       ScoreUser       `json:"user"`
}
