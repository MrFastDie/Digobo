package twitch

import (
	"Digobo/config"
	"Digobo/log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Infos struct {
	Data []struct {
		BroadcasterId       string   `json:"broadcaster_id"`
		BroadcasterLogin    string   `json:"broadcaster_login"`
		BroadcasterName     string   `json:"broadcaster_name"`
		BroadcasterLanguage string   `json:"broadcaster_language"`
		GameId              string   `json:"game_id"`
		GameName            string   `json:"game_name"`
		Title               string   `json:"title"`
		Delay               int      `json:"delay"`
		Tags                []string `json:"tags"`
	} `json:"data"`
}

func GetInfos(userIds ...string) Infos {
	url := "https://api.twitch.tv/helix/channels"

	broadcasterPayload := "?broadcaster_id=" + strings.Join(userIds, "&broadcaster_id=")
	req, err := http.NewRequest("GET", url+broadcasterPayload, nil)
	if err != nil {
		log.Error.Println(err)
	}

	token, err := getToken()
	if err != nil {
		log.Error.Println("OAuth key unavailable - skipping")
		return Infos{}
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Client-ID", config.Config.Apps.Twitch.ClientId)
	respRaw, err := client.Do(req)
	if err != nil {
		log.Error.Println(err)
		return Infos{}
	}

	defer respRaw.Body.Close()
	respBytes, err := ioutil.ReadAll(respRaw.Body)
	if err != nil {
		log.Error.Println(err)
	}

	var infos Infos
	err = json.Unmarshal(respBytes, &infos)

	return infos
}
