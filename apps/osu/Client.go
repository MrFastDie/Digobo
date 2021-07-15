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
const SCORE_URL = "/api/v2/users/%d/scores/%s?mode=%s&limit=%d&offset=%d"

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

func GetUserRecentActivity(userId int) (UserRecentActivityResult, error) {
	var ret UserRecentActivityResult

	err := requestToModel(prepareRequest("GET", fmt.Sprintf(USER_RECENT_ACTIVITY_URL, strconv.Itoa(userId)), getToken()), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func GetUser(userId int) (UserProfile, error) {
	var ret UserProfile

	err := requestToModel(prepareRequest("GET", fmt.Sprintf(OSU_API_URL + USER_PROFILE_URL, userId), getToken()), &ret)
	if err != nil {
		return UserProfile{}, err
	}

	return ret, nil
}

func GetScores(userId int, scoreType ScoreType, scoreMode ScoreMode, limit int, offset int) ([]Score, error) {
	var ret []Score

	req := prepareRequest("GET", fmt.Sprintf(OSU_API_URL + SCORE_URL, userId, scoreType, scoreMode, limit, offset), getToken())
	err := requestToModel(req, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func prepareRequest(method string, url string, token string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)

	return req
}

func requestToModel(req *http.Request, model interface{}) error {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Error.Println("can't do osu! request", err)
		return err
	}

	err = json.Json.NewDecoder(res.Body).Decode(model)
	if err != nil {
		var errStruct struct {
			Authentication string `json:"authentication"`
		}
		newErr := json.Json.NewDecoder(res.Body).Decode(errStruct)
		if newErr != nil {
			log.Error.Println("can't unmarshal to requested model", err)
			return err
		}

		// Token was invalid, generating new one
		generateToken()

		return nil
	}

	return nil
}

func getToken() string {
	if token.TokenType == "" || token.ExpiresAt.Sub(time.Now()) <= 0 {
		generateToken()
	}

	return token.AccessToken
}

func generateToken() {
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
		return
	}

	err = json.Json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		log.Error.Println("can't unmarshal token for osu!", err)
		return
	}

	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
}