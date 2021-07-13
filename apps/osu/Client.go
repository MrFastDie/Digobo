package osu

import (
	"Digobo/config"
	"Digobo/json"
	"Digobo/log"
	"fmt"
	"net/http"
	"strconv"
)

const OSU_API_URL = "https://osu.ppy.sh"
const USER_RECENT_ACTIVITY_URL = OSU_API_URL + "/api/v2/users/%s/recent_activity"

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

func GetUserRecentActivity(userId int) (UserRecentActivityResult, error) {
	var ret UserRecentActivityResult
	client := &http.Client{}

	res, err := client.Do(prepareRequest("GET", fmt.Sprintf(USER_RECENT_ACTIVITY_URL, strconv.Itoa(userId))))
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

func GetUserBeatmaps(userId int, beatmapType BeatmapType) UserBeatmapsResult {
	var ret UserBeatmapsResult

	client := &http.Client{}

	if USER_BEATMAPS_FAVORITE&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" +USER_BEATMAPS_STRING_FAVORITE+ "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Favorite)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu USER_BEATMAPS_FAVORITE", err)
		}
	}

	if USER_BEATMAPS_GRAVEYARD&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" +USER_BEATMAPS_STRING_GRAVEYARD+ "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Graveyard)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu USER_BEATMAPS_GRAVEYARD", err)
		}
	}

	if USER_BEATMAPS_LOVED&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" +USER_BEATMAPS_STRING_LOVED+ "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Loved)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu loved", err)
		}
	}

	if USER_BEATMAPS_MOST_PLAYED&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" +USER_BEATMAPS_STRING_MOST_PLAYED+ "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.MostPlayer)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu USER_BEATMAPS_MOST_PLAYED", err)
		}
	}

	if USER_BEATMAPS_PENDING&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" +USER_BEATMAPS_STRING_PENDING+ "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Pending)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu USER_BEATMAPS_PENDING", err)
		}
	}

	if USER_BEATMAPS_RANKED&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" +USER_BEATMAPS_STRING_RANKED+ "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Ranked)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu USER_BEATMAPS_RANKED", err)
		}
	}

	return ret
}

func prepareRequest(method string, url string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Config.Apps.Osu.Token)

	return req
}