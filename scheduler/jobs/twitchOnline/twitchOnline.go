package twitchOnline

import (
	"Digobo/apps/twitch"
	"Digobo/config"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/log"
	"Digobo/scheduler"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

type Data struct {
	UserId    string
	ChannelId string
}

type TwitchOnline struct{}

func (this *TwitchOnline) Execute(rawData string) error {
	var jobData Data

	err := json.Unmarshal([]byte(rawData), &jobData)
	if err != nil {
		log.Error.Println("Cant unmarshal data for scheduled job", err)
		return err
	}

	onlineData := twitch.GetStreams(jobData.UserId)

	dbData, err := database.GetTwitchWatcherByUserAndChannel(jobData.UserId, jobData.ChannelId)
	if err != nil {
		return err
	}

	if 0 != len(onlineData.Data) && !dbData.Online {
		thumbnail := strings.Replace(strings.Replace(onlineData.Data[0].ThumbnailUrl, "{width}", "350", 1), "{height}", "200", 1)
		data := onlineData.Data[0]

		username := data.UserName

		mature := "False"
		if data.IsMature {
			mature = "True"
		}

		embed := &discordgo.MessageEmbed{
			Title:       username + " is live!",
			Description: data.Title,
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       config.Config.Bot.DefaultEmbedColor,
			Author:      &discordgo.MessageEmbedAuthor{},
			Image: &discordgo.MessageEmbedImage{
				URL:    thumbnail,
				Width:  350,
				Height: 200,
			},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Language",
					Value:  data.Language,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Mature Audiances",
					Value:  mature,
					Inline: true,
				},
			},
		}

		if data.GameName != "" {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Playing",
				Value:  data.GameName,
				Inline: true,
			})
		}

		_, err := discordBot.GetInstance().ChannelMessageSendEmbed(jobData.ChannelId, embed)
		if err != nil {
			log.Warning.Println("couldn't send embed to discord server channel for notify", err)
		}

		err = database.ChangeTwitchWatcherStatusByUserAndChannel(jobData.UserId, jobData.ChannelId, true)
		if err != nil {
			return err
		}
	} else if 0 == len(onlineData.Data) && dbData.Online {
		err = database.ChangeTwitchWatcherStatusByUserAndChannel(jobData.UserId, jobData.ChannelId, false)
		if err != nil {
			return err
		}
	}

	job := scheduler.Job{
		ExecutionTime: time.Now().Add(2 * time.Minute),
		ExecutionFunc: &TwitchOnline{},
		Data:          rawData,
	}

	scheduler.GetScheduler().AddThresholdJob(job)

	return nil
}
