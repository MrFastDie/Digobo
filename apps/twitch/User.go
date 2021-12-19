package twitch

import (
	"Digobo/config"
	"Digobo/json"
	"Digobo/log"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Data []struct {
		Id              string    `json:"id"`
		Login           string    `json:"login"`
		DisplayName     string    `json:"display_name"`
		Type            string    `json:"type"`
		BroadcasterType string    `json:"broadcaster_type"`
		Description     string    `json:"description"`
		ProfileImageUrl string    `json:"profile_image_url"`
		OfflineImageUrl string    `json:"offline_image_url"`
		ViewCount       int       `json:"view_count"`
		CreatedAt       time.Time `json:"created_at"`
	} `json:"data"`
}

func GetUserByLogin(loginName string) User {
	url := "https://api.twitch.tv/helix/users"

	broadcasterPayload := "?login=" + loginName
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

	var user User
	err = json.Unmarshal(respBytes, &user)

	return user
}
