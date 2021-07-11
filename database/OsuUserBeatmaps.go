package database

import (
	"Digobo/apps/osu"
	"Digobo/json"
)

type OsuUserBeatmaps struct {
	Uuid        string                 `db:"uuid"`
	User        int                    `db:"user"`
	UserName    string                 `db:"user_name"`
	BeatmapType osu.BeatmapType        `db:"beatmap_type"`
	BeatmapData osu.UserBeatmapsResult `db:"beatmap_data"`
}

func (this *OsuUserBeatmaps) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

func GetOsuUserBeatmaps(userId int) (OsuUserBeatmaps, error) {
	var ret OsuUserBeatmaps
	err := db.Get(&ret, `SELECT row_to_json(t) FROM (SELECT * FROM osu_user_beatmaps WHERE "user" = $1) t;`, userId)

	if err != nil {
		return OsuUserBeatmaps{}, err
	}

	return ret, nil
}

func InsertIntoOsuUserBeatmaps(userId int, beatmapType osu.BeatmapType, beatmapData osu.UserBeatmapsResult) error {
	rawData, _ := json.Json.Marshal(beatmapData)
	_, err := db.Exec(`INSERT INTO osu_user_beatmaps ("user", beatmap_type, beatmap_data) VALUES ($1, $2, $3)`, userId, beatmapType, string(rawData))
	if err != nil {
		return err
	}

	return nil
}

func UpdateOsuUserBeatmaps(uuid string, beatmapType osu.BeatmapType, beatmapData osu.UserBeatmapsResult) error {
	rawData, _ := json.Json.Marshal(beatmapData)
	_, err := db.Exec(`UPDATE osu_user_beatmaps SET beatmap_type = $1, beatmap_data = $2 WHERE uuid=$3`, beatmapType, rawData, uuid)
	if err != nil {
		return err
	}

	return nil
}
