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
	"github.com/spf13/cobra"
	"time"
)

var listOsuUserWatcher = &cobra.Command{
	Use:   "list",
	Short: "list all users from your personal watch list",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		i := command.CommandI

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

		err = discordBot.SendEmbed(embed, i.ChannelID, s)
		if err != nil {
			return err
		}
		return nil
	},
}
