package twitch

import (
	"Digobo/config"
	"Digobo/log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	TokenType   string    `json:"token_type"`
	ExpiresAt   time.Time `json:"-"`
}

var savedToken *Token

func getToken() (*Token, error) {
	if nil == savedToken || savedToken.ExpiresAt.After(time.Now()) {
		var url = "https://id.twitch.tv/oauth2/token"
		var payload, _ = json.Marshal(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     config.Config.Apps.Twitch.ClientId,
			"client_secret": config.Config.Apps.Twitch.ClientSecret,
		})

		req := Token{}

		requestBody := bytes.NewBuffer(payload)

		respRaw, err := http.Post(url, "application/json", requestBody)
		if err != nil {
			log.Error.Println(err)
			return nil, err
		}

		defer respRaw.Body.Close()
		respBytes, err := ioutil.ReadAll(respRaw.Body)
		if err != nil {
			log.Error.Println(err)
			return nil, err
		}

		err = json.Unmarshal(respBytes, &req)
		if err != nil {
			log.Error.Println(err)
			return nil, err
		}

		req.ExpiresAt = time.Now().Add(time.Duration(req.ExpiresIn) * time.Second)
		savedToken = &req
	}

	return savedToken, nil
}
