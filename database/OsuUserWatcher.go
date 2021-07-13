package database

import "Digobo/json"

type OsuUserWatcher struct {
	UserId        int      `db:"user_id"`
	UserName      string   `db:"user_name"`
	OutputChannel []string `db:"output_channel"`
}

func (this *OsuUserWatcher) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

func GetOsuWatcher() ([]OsuUserWatcher, error) {
	var watcher []OsuUserWatcher

	err := db.Select(&watcher, `SELECT row_to_json(d)
FROM (SELECT osu_user_watcher.user_id,
             user_name,
             (
                 SELECT array_agg(channel_id)
                 FROM osu_user_watcher_channel
                 WHERE osu_user_watcher.user_id = osu_user_watcher_channel.user_id
             ) as output_channel
      FROM osu_user_watcher) d`)
	if err != nil {
		return nil, err
	}

	return watcher, nil
}