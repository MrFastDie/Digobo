package notifyEventParticipants

import (
	"Digobo/config"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/json"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"time"
)

type Data struct {
	ParentEventUuid string                        `json:"parent_event_uuid"`
	NotifyChannels  []string                      `json:"notify_channels"`
	NotifyUsers     []string                      `json:"notify_users"`
	OriginDate      json.TimestampWithoutTimezone `json:"origin_time"`
}

type NotifyEventParticipants struct{}

func (this *NotifyEventParticipants) Execute(data string) error {
	var notifyData Data
	err := json.Default.Unmarshal([]byte(data), &notifyData)
	if err != nil {
		log.Error.Println("can't unmarshal notifyEventData to struct", err)
		return err
	}

	parentEvent, err := database.GetEventByUuid(notifyData.ParentEventUuid)
	if err != nil {
		return err
	}

	var dinstance = discordBot.GetInstance()

	parentEventStr, _ := json.Default.Marshal(parentEvent)

	embed := &discordgo.MessageEmbed{
		Title:       parentEvent.Title,
		Description: parentEvent.Description,
		Timestamp:   notifyData.OriginDate.Time.Format(time.RFC3339),
		Color:       config.Config.Bot.DefaultEmbedColor,
		Author:      &discordgo.MessageEmbedAuthor{},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Debugdata", // TODO remove
				Value:  string(parentEventStr),
				Inline: true,
			},
		},
	}

	for _, notifyUser := range notifyData.NotifyUsers {
		channel, err := dinstance.UserChannelCreate(notifyUser)
		if err != nil {
			log.Warning.Println("can't create channel with discord user to notify", err)
			continue
		}

		_, err = dinstance.ChannelMessageSendEmbed(channel.ID, embed)
		if err != nil {
			log.Warning.Println("couldn't send embed to discord user for notify", err)
			continue
		}
	}

	for _, notifyChannel := range notifyData.NotifyChannels {
		_, err := dinstance.ChannelMessageSendEmbed(notifyChannel, embed)
		if err != nil {
			log.Warning.Println("couldn't send embed to discord server channel for notify", err)
			continue
		}
	}

	return nil
}
