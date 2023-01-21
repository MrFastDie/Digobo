package osuUserWatcher

import (
	"Digobo/apps/osu"
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
	Name:        "osu",
	Description: "List all watched Osu! user",
	SubCommands: map[string]command.SubCommand{
		Add.Name:    Add,
		Color.Name:  Color,
		Remove.Name: Remove,
	},
	Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
		discordBot.SendInteractionMessage("Command received", s, i)

		list, err := database.GetOsuWatcherListByChannel(i.ChannelID)
		if err != nil && err.Error() == database.NO_ROWS {
			discordBot.SendMessage("There are currently no member in your watch list", i.ChannelID, s)
			return nil
		} else if err != nil {
			log.Error.Println("can't fetch osu watcher list", err)
			return err
		}

		var fields []*discordgo.MessageEmbedField

		for _, entry := range list {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   fmt.Sprintf("%s (%d)", entry.UserName, entry.UserId),
				Value:  fmt.Sprintf("%s/users/%d", osu.OSU_API_URL, entry.UserId),
				Inline: false,
			})
		}

		embed := &discordgo.MessageEmbed{
			Title:       "osu! watcher list",
			Description: "Here is everyone currently listed by this channel",
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       config.Config.Bot.DefaultEmbedColor,
			Fields:      fields,
		}

		return discordBot.SendEmbed(embed, i.ChannelID, s)
	},
}

func hasChannel(channelId string, channelList []database.OsuOutputChannel) bool {
	for _, b := range channelList {
		if b.ChannelId == channelId {
			return true
		}
	}
	return false
}

func init() {
	command.Map[Command.Name] = Command
}
