package twitch

import (
	"Digobo/config"
	"Digobo/log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Streams struct {
	Data []struct {
		Id           string    `json:"id"`
		UserId       string    `json:"user_id"`
		UserLogin    string    `json:"user_login"`
		UserName     string    `json:"user_name"`
		GameId       string    `json:"game_id"`
		GameName     string    `json:"game_name"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailUrl string    `json:"thumbnail_url"`
		TagIds       []string  `json:"tag_ids"`
		IsMature     bool      `json:"is_mature"`
	} `json:"data"`
	Pagination struct {
	} `json:"pagination"`
}

func GetStreams(userIds ...string) Streams {
	url := "https://api.twitch.tv/helix/streams"

	broadcasterPayload := "?user_id=" + strings.Join(userIds, "&user_id=")
	req, err := http.NewRequest("GET", url+broadcasterPayload, nil)
	if err != nil {
		log.Error.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+getToken().AccessToken)
	req.Header.Set("Client-ID", config.Config.Apps.Twitch.ClientId)
	respRaw, err := client.Do(req)
	if err != nil {
		log.Error.Println(err)
	}

	defer respRaw.Body.Close()
	respBytes, err := ioutil.ReadAll(respRaw.Body)
	if err != nil {
		log.Error.Println(err)
	}

	var streams Streams
	err = json.Unmarshal(respBytes, &streams)

	return streams
}

func IsOnline(userId string) bool {
	return 1 == len(GetStreams(userId).Data)
}
