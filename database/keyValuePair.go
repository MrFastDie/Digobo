package database

import "Digobo/json"

type KeyValuePair struct {
	Uuid      string `db:"uuid" json:"uuid"`
	Command   string `db:"command" json:"command"`
	Value     string `db:"value" json:"value"`
	DiscordID string `db:"discord_id" json:"discord_id"`
}

func (this *KeyValuePair) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

func GetKeyValuePairByCommand(command string) (string, error) {
	var ret string

	err := db.Get(&ret, `SELECT value FROM random_answer_list WHERE command = $1 ORDER BY random() limit 1`, command)
	if err != nil {
		return "", err
	}

	return ret, nil
}

func AddKeyValuePairByCommand(command string, value string, discordId string) error {
	_, err := db.Exec(`INSERT INTO random_answer_list (command, value, discord_id) VALUES ($1, $2, $3)`, command, value, discordId)

	return err
}

func RemoveKeyValuePairByCommandAndValue(command string, value string) error {
	_, err := db.Exec(`DELETE FROM random_answer_list WHERE command = $1 AND value = $2`, command, value)

	return err
}
