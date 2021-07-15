package osu

import (
	"Digobo/config"
	"Digobo/json"
	"Digobo/log"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const OSU_API_URL = "https://osu.ppy.sh"
const OSU_TOKEN_URL = "/oauth/token"
const USER_RECENT_ACTIVITY_URL = OSU_API_URL + "/api/v2/users/%s/recent_activity"
const USER_PROFILE_URL = "/api/v2/users/%d"

type BeatmapType int

type UserBeatmapsResult struct {
	Favorite   []UserBeatmaps
	Graveyard  []UserBeatmaps
	Loved      []UserBeatmaps
	MostPlayer []UserBeatmaps
	Pending    []UserBeatmaps
	Ranked     []UserBeatmaps
}

type UserRecentActivityResult []Event

type OAuthTokenRequest struct {
	ClientId     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type OAuthToken struct {
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"-"`
	AccessToken string    `json:"access_token"`
}

var token OAuthToken

func getToken() string {
	// TODO validate token func if a result cant be parsed
	if token.TokenType == "" || token.ExpiresAt.Sub(time.Now()) <= 0 {
		tokenReq := OAuthTokenRequest{
			ClientId: config.Config.Apps.Osu.ClientId,
			ClientSecret: config.Config.Apps.Osu.ClientSecret,
			GrantType: "client_credentials",
			Scope: "public",
		}

		tokenReqStr, _ := json.Json.Marshal(tokenReq)

		req, _ := http.NewRequest("POST", OSU_API_URL + OSU_TOKEN_URL, bytes.NewBuffer(tokenReqStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			log.Error.Println("can't fetch new token for osu!", err)
			return ""
		}

		err = json.Json.NewDecoder(res.Body).Decode(&token)
		if err != nil {
			log.Error.Println("can't unmarshal token for osu!", err)
			return ""
		}

		token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	return token.AccessToken
}

func GetUserRecentActivity(userId int) (UserRecentActivityResult, error) {
	var ret UserRecentActivityResult
	client := &http.Client{}

	res, err := client.Do(prepareRequest("GET", fmt.Sprintf(USER_RECENT_ACTIVITY_URL, strconv.Itoa(userId)), getToken()))
	if err != nil {
		log.Error.Println("can't fetch from osu! api", err)
		return nil, err
	}

	err = json.Json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		log.Warning.Println("can't decode osu! api result to event object", err)
		return nil, err
	}

	return ret, nil
}

func GetUser(userId int) (UserProfile, error) {
	var ret UserProfile
	client := &http.Client{}

	req := prepareRequest("GET", fmt.Sprintf(OSU_API_URL + USER_PROFILE_URL, userId), getToken())

	res, err := client.Do(req)
	if err != nil {
		log.Error.Println("can't fetch user info from osu!", err)
		return UserProfile{}, err
	}

	err = json.Json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		log.Error.Println("can't unmarshal osu! user api data to model", err)
		return UserProfile{}, err
	}

	return ret, nil
}

func prepareRequest(method string, url string, token string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)

	return req
}
