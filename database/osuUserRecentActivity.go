package database

import "Digobo/json"

type OsuUserRecentActivity struct {
	Uuid           string `db:"uuid"`
	UserId         int    `db:"user_id"`
	LastActivityId int    `db:"last_activity_id"`
}

func (this *OsuUserRecentActivity) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

func OsuUserRecentActivityGetLastActivityId(userId int) (int, error) {
	var recentActivity OsuUserRecentActivity

	err := db.Get(&recentActivity, `SELECT row_to_json(t) FROM (SELECT * FROM osu_user_recent_activity WHERE user_id = $1) t;`, userId)
	if err != nil {
		return 0, err
	}

	return recentActivity.LastActivityId, nil
}

func OsuUserRecentActivityUpdateLastActivityId(userId int, activityId int) error {
	_, err := db.Exec(`UPDATE osu_user_recent_activity SET last_activity_id=$1 WHERE user_id = $2`, activityId, userId)
	if err != nil {
		return err
	}

	return nil
}

func OsuUserRecentActivityInsertLastActivity(userId int, activityId int) error {
	_, err := db.Exec(`INSERT INTO osu_user_recent_activity (user_id, last_activity_id) VALUES ($1, $2)`, userId, activityId)
	if err != nil {
		return err
	}

	return nil
}
