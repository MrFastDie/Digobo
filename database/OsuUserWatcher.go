package database

import "Digobo/json"

type OsuOutputChannel struct {
	ChannelId string `db:"channel_id"`
	Color     *int   `db:"color"`
}

type OsuUserWatcher struct {
	UserId        int                `db:"user_id"`
	UserName      string             `db:"user_name"`
	OutputChannel []OsuOutputChannel `db:"output_channel"`
}

func (this *OsuUserWatcher) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

type OsuWatcherList struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"user_name"`
}

func (this *OsuWatcherList) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

func GetOsuWatchers() ([]OsuUserWatcher, error) {
	var watcher []OsuUserWatcher

	err := db.Select(&watcher, `SELECT row_to_json(d)
FROM (SELECT osu_user_watcher.user_id,
             user_name,
             (
                 SELECT array_agg(row_to_json(d)) FROM (SELECT channel_id, color
                 FROM osu_user_watcher_channel
                 WHERE osu_user_watcher.user_id = osu_user_watcher_channel.user_id
             )d) as output_channel
      FROM osu_user_watcher) d`)
	if err != nil {
		return nil, err
	}

	return watcher, nil
}

func GetOsuWatcher(userId int) (OsuUserWatcher, error) {
	var watcher OsuUserWatcher

	err := db.Get(&watcher, `SELECT row_to_json(d)
FROM (SELECT osu_user_watcher.user_id,
             user_name,
             (
                 SELECT array_agg(row_to_json(d)) FROM (SELECT channel_id, color
                 FROM osu_user_watcher_channel
                 WHERE osu_user_watcher.user_id = osu_user_watcher_channel.user_id
             )d) as output_channel
      FROM osu_user_watcher WHERE user_id=$1) d`, userId)
	if err != nil {
		return OsuUserWatcher{}, err
	}

	return watcher, nil
}

func AddOsuWatcherOutputChannel(userId int, channelId string) error {
	_, err := db.Exec(`INSERT INTO osu_user_watcher_channel (user_id, channel_id) VALUES ($1, $2)`, userId, channelId)
	if err != nil {
		return err
	}

	return nil
}

func AddOsuWatcherUser(userId int, userName string) error {
	_, err := db.Exec(`INSERT INTO osu_user_watcher (user_id, user_name) VALUES ($1, $2)`, userId, userName)
	if err != nil {
		return err
	}

	return nil
}

func RemoveOsuWatcherOutputChannel(userId int, channelId string) error {
	_, err := db.Exec(`DELETE FROM osu_user_watcher_channel WHERE user_id=$1 AND channel_id=$2`, userId, channelId)
	if err != nil {
		return err
	}

	return nil
}

func AddOsuWatcherColor(userId int, channelId string, color int) error {
	_, err := db.Exec(`UPDATE osu_user_watcher_channel SET color=$1 WHERE user_id=$2 AND channel_id=$3`, color, userId, channelId)
	if err != nil {
		return err
	}

	return nil
}

func GetOsuWatcherListByChannel(channelId string) ([]OsuWatcherList, error) {
	var list []OsuWatcherList

	err := db.Select(&list, `SELECT row_to_json(d)
FROM (SELECT osu_user_watcher_channel.user_id, user_name
      FROM osu_user_watcher_channel,
           osu_user_watcher
      WHERE osu_user_watcher_channel.user_id = osu_user_watcher.user_id
        AND channel_id = $1) d`, channelId)
	if err != nil {
		return nil, err
	}

	return list, nil
}