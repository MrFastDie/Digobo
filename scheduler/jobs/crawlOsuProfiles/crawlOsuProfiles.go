package CrawlOsuProfiles

import (
	"Digobo/apps/osu"
	"Digobo/config"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/json"
	"Digobo/log"
	"Digobo/scheduler"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

type Data struct {
	UserProfile  int
	BeatmapTypes osu.BeatmapType
}

type CrawlOsuProfiles struct{}

func (this *CrawlOsuProfiles) Execute(rawData string) error {
	var data Data

	if rawData == "" {
		data.UserProfile = 13467065
		data.BeatmapTypes = osu.GRAVEYARD | osu.PENDING | osu.RANKED | osu.LOVED
	} else {
		json.Unmarshal(rawData, data)
	}

	profileResult := osu.GetUserBeatmaps(data.UserProfile, data.BeatmapTypes)
	databaseResult, err := database.GetOsuUserBeatmaps(data.UserProfile)
	if err != nil && err.Error() == "sql: no rows in result set" {
		err := database.InsertIntoOsuUserBeatmaps(data.UserProfile, data.BeatmapTypes, profileResult)
		if err != nil {
			log.Warning.Println(err)
		}
	} else if err != nil {
		log.Warning.Println(err)
	} else {
		// TODO only show the really new ones (or deleted)
		// TODO better comparement len != len is kinda cheap but deepequal wont work
		if len(profileResult.Graveyard) != len(databaseResult.BeatmapData.Graveyard) {
			notifyChannel(databaseResult, osu.STRING_GRAVEYARD, profileResult.Graveyard)
		}

		if len(profileResult.Pending) != len(databaseResult.BeatmapData.Pending) {
			notifyChannel(databaseResult, osu.STRING_PENDING, profileResult.Pending)
		}

		if len(profileResult.Ranked) != len(databaseResult.BeatmapData.Ranked) {
			notifyChannel(databaseResult, osu.STRING_RANKED, profileResult.Ranked)
		}

		if len(profileResult.Loved) != len(databaseResult.BeatmapData.Loved) {
			notifyChannel(databaseResult, osu.STRING_LOVED, profileResult.Loved)
		}

		err := database.UpdateOsuUserBeatmaps(databaseResult.Uuid, databaseResult.BeatmapType, profileResult)
		if err != nil {
			log.Warning.Println(err)
		}
	}

	dataStr, _ := json.Json.Marshal(data)

	CrawlOsuProfilesJobStart(time.Now().Add(2 * time.Hour), string(dataStr))

	return nil
}

func notifyChannel(dbData database.OsuUserBeatmaps, beatmapType string, data []osu.UserBeatmaps) {
	var embedData []*discordgo.MessageEmbedField

	for _, beatmap := range data {
		embedData = append(embedData, &discordgo.MessageEmbedField{
			Name:   beatmap.Title,
			Value:  "https://osu.ppy.sh/beatmapsets/" + strconv.Itoa(beatmap.Id),
			Inline: false,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:       dbData.UserName,
		Description: beatmapType + " has changed content",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       config.Config.Bot.DefaultEmbedColor,
		Author:      &discordgo.MessageEmbedAuthor{},
		Fields: embedData,
	}

	embed.URL = "https://osu.ppy.sh/users/" + strconv.Itoa(dbData.User)

	// TODO send to specific server channel
	_, err := discordBot.GetInstance().ChannelMessageSendEmbed("863777071113043989", embed)
	if err != nil {
		log.Warning.Println("couldn't send embed to discord server channel for notify", err)
	}
}

func CrawlOsuProfilesJobStart(runTime time.Time, data string) {
	job := scheduler.Job{
		ExecutionTime: runTime,
		ExecutionFunc: &CrawlOsuProfiles{},
		Data:          data,
	}

	scheduler.GetScheduler().AddThresholdJob(job)
}
