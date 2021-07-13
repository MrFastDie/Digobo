package osu

import (
	"Digobo/config"
	"Digobo/json"
	"Digobo/log"
	"net/http"
	"strconv"
)

const OSU_API_URL = "https://osu.ppy.sh"

type BeatmapType int

type UserBeatmapsResult struct {
	Favorite   []UserBeatmaps
	Graveyard  []UserBeatmaps
	Loved      []UserBeatmaps
	MostPlayer []UserBeatmaps
	Pending    []UserBeatmaps
	Ranked     []UserBeatmaps
}

const (
	FAVORITE    BeatmapType = 1 << iota
	GRAVEYARD   BeatmapType = 1 << iota
	LOVED       BeatmapType = 1 << iota
	MOST_PLAYED BeatmapType = 1 << iota
	PENDING     BeatmapType = 1 << iota
	RANKED      BeatmapType = 1 << iota
)

const (
	STRING_FAVORITE = "favorite"
	STRING_GRAVEYARD = "graveyard"
	STRING_LOVED = "loved"
	STRING_MOST_PLAYED = "most_played"
	STRING_PENDING = "pending"
	STRING_RANKED = "ranked"
)

func prepareRequest(method string, url string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Config.Apps.Osu.Token)

	return req
}

func GetUserBeatmaps(userId int, beatmapType BeatmapType) UserBeatmapsResult {
	var ret UserBeatmapsResult

	client := &http.Client{}

	if FAVORITE&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" + STRING_FAVORITE + "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Favorite)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu FAVORITE", err)
		}
	}

	if GRAVEYARD&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" + STRING_GRAVEYARD + "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Graveyard)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu GRAVEYARD", err)
		}
	}

	if LOVED&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" + STRING_LOVED + "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Loved)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu loved", err)
		}
	}

	if MOST_PLAYED&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" + STRING_MOST_PLAYED + "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.MostPlayer)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu MOST_PLAYED", err)
		}
	}

	if PENDING&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" + STRING_PENDING + "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Pending)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu PENDING", err)
		}
	}

	if RANKED&beatmapType > 0 {
		res, _ := client.Do(prepareRequest("GET", OSU_API_URL+"/api/v2/users/"+strconv.Itoa(userId)+"/beatmapsets/" + STRING_RANKED + "?limit=2000"))

		err := json.Json.NewDecoder(res.Body).Decode(&ret.Ranked)
		if err != nil {
			log.Warning.Println("Cant unmarshal osu RANKED", err)
		}
	}

	return ret
}
