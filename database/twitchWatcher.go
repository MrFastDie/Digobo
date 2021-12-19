package database

import "Digobo/json"

type TwitchWatcher struct {
	Uuid      string `db:"uuid"`
	UserId    string `db:"user_id"`
	ChannelId string `db:"channel_id"`
	Online    bool   `db:"online"`
}

func (this *TwitchWatcher) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

// TODO rework that twitch user only need to get fetched once instead of for every channel

func GetTwitchWatcher() ([]TwitchWatcher, error) {
	var ret []TwitchWatcher
	err := db.Select(&ret, `SELECT row_to_json(d) FROM (SELECT * FROM twitch_watcher) d`)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func GetTwitchWatcherByUserAndChannel(userId, channelId string) (TwitchWatcher, error) {
	var ret TwitchWatcher
	err := db.Get(&ret, `SELECT row_to_json(d) FROM (SELECT * FROM twitch_watcher WHERE user_id = $1 AND channel_id = $2) d`, userId, channelId)
	if err != nil {
		return TwitchWatcher{}, err
	}

	return ret, nil
}

func AddTwitchWatcher(userId, channelId string, online bool) error {
	_, err := db.Exec(`INSERT INTO twitch_watcher (user_id, channel_id, online) VALUES ($1, $2, $3)`, userId, channelId, online)
	if err != nil {
		return err
	}

	return nil
}

func ChangeTwitchWatcherStatusByUserAndChannel(userId, channelId string, status bool) error {
	_, err := db.Exec(`UPDATE twitch_watcher SET online=$1 WHERE user_id = $2 and channel_id = $3`, status, userId, channelId)
	if err != nil {
		return err
	}

	return nil
}
