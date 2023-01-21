package twitch

import (
	twitch2 "Digobo/apps/twitch"
	"Digobo/config"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

var Command = command.Command{
	Name:        "twitch",
	Description: "List all watched twitch channel",
	SubCommands: map[string]command.SubCommand{
		AddChannel.Name: AddChannel,
	},
	Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
		discordBot.SendInteractionMessage("Command received", s, i)

		list, err := database.GetTwitchWatchersByChannel(i.ChannelID)
		if err != nil && err.Error() == database.NO_ROWS {
			discordBot.SendMessage("There are currently no member in your watch list", i.ChannelID, s)
			return nil
		} else if err != nil {
			log.Error.Println("can't fetch twitch watcher list", err)
			return err
		}

		var fields []*discordgo.MessageEmbedField
		var streamNames []string

		for _, entry := range list {
			streamNames = append(streamNames, entry.UserId)
		}

		onlineData := twitch2.GetInfos(streamNames...)
		for _, data := range onlineData.Data {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   fmt.Sprintf("%s", data.BroadcasterName),
				Value:  fmt.Sprintf("[Visit %s](https://twitch.tv/%s)\nCurrently playing: %s\nTitle: %s", data.BroadcasterName, data.BroadcasterName, data.GameName, data.Title),
				Inline: false,
			})
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Twitch watcher list",
			Description: "Here is everyone currently listed by this channel\nPlease note: This does not indicate if someone is live!",
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       config.Config.Bot.DefaultEmbedColor,
			Fields:      fields,
		}

		return discordBot.SendEmbed(embed, i.ChannelID, s)
	},
}

func init() {
	command.Map[Command.Name] = Command
}
