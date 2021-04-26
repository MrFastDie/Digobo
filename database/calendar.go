package database

import "Digobo/json"

type Calendar struct {
	Uuid             string `db:"uuid" json:"uuid"`
	Name             string `db:"name" json:"name"`
	Description      string `db:"description" json:"description"`
	CreatorDiscordId string `db:"creator_discord_id" json:"creator_discord_id"`
}

func (this *Calendar) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}
