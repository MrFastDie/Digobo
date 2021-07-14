package CrawlOsuProfiles

import (
	"Digobo/apps/osu"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/json"
	"Digobo/log"
	"Digobo/scheduler"
	"github.com/bwmarrin/discordgo"
	"time"
)

type Data struct {
	UserId        int
	UserName      string
	OutputChannel []database.OsuOutputChannel
}

type CrawlOsuProfiles struct{}

func (this *CrawlOsuProfiles) Execute(rawData string) error {
	var data Data

	err := json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		log.Error.Println("Cant unmarshal data for scheduled job", err)
		return err
	}

	osuUserDbData, err := database.GetOsuWatcher(data.UserId)
	if err != nil {
		log.Error.Println("can't fetch osu! DB user!", err)
		return err
	}

	data.OutputChannel = osuUserDbData.OutputChannel
	if len(data.OutputChannel) == 0 {
		return nil
	}

	var osuData osu.UserRecentActivityResult

	recentActivityId, err := database.OsuUserRecentActivityGetLastActivityId(data.UserId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}

	osuData, err = osu.GetUserRecentActivity(data.UserId)
	if err != nil {
		return err
	}

	var notifyData []*discordgo.MessageEmbed

	for _, osuDataSet := range osuData {
		if recentActivityId == osuDataSet.Id {
			break
		}

		switch osuDataSet.Type {
		case osu.EVENT_TYPE_ACHIEVEMENT,
			osu.EVENT_TYPE_BEATMAP_PLAYCOUNT,
			osu.EVENT_TYPE_BEATMAPSET_APPROVE,
			osu.EVENT_TYPE_BEATMAPSET_DELETE,
			osu.EVENT_TYPE_BEATMAPSET_REVIVE,
			osu.EVENT_TYPE_BEATMAPSET_UPDATE,
			osu.EVENT_TYPE_BEATMAPSET_UPLOAD,
			osu.EVENT_TYPE_RANK,
			osu.EVENT_TYPE_RANK_LOST:
			notifyData = append(notifyData, osuDataSet.PrepareEmbed(data.UserName, data.UserId))
			break
		}
	}

	for i := len(notifyData)-1; i >= 0; i-- {
		for _, channel := range data.OutputChannel {
			notifyChannel(channel.ChannelId, notifyData[i], channel.Color)
		}
	}

	if len(osuData) > 0 {
		if recentActivityId == 0 {
			err := database.OsuUserRecentActivityInsertLastActivity(data.UserId, osuData[0].Id)
			if err != nil {
				log.Error.Println("can't insert into osu_user_recent_activity", err)
				return err
			}
		} else {
			err := database.OsuUserRecentActivityUpdateLastActivityId(data.UserId, osuData[0].Id)
			if err != nil {
				log.Error.Println("can't update osu_user_recent_activity", err)
				return err
			}
		}
	}

	CrawlOsuProfilesJobStart(time.Now().Add(5*time.Minute), rawData)

	return nil
}

func notifyChannel(channelId string, embed *discordgo.MessageEmbed, color *int) {
	if color != nil {
		embed.Color = *color
	}

	_, err := discordBot.GetInstance().ChannelMessageSendEmbed(channelId, embed)
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
