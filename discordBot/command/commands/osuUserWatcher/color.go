package osuUserWatcher

import (
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"github.com/spf13/cobra"
	"strconv"
)

var colorOsuUserWatcher = &cobra.Command{
	Use:   "color [user_id] [color_in_hex_without_#]",
	Short: "changes the embed color of the given user id",
	Long:  "This command allows you to change the appearance color of the embed created when the user has new updates",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		m := command.CommandM

		userId, err := strconv.Atoi(args[0])
		if err != nil {
			discordBot.SendMessage("Please provide a valid user_id", m.ChannelID, s)
		}

		user, err := database.GetOsuWatcher(userId)
		if err != nil || !hasChannel(m.ChannelID, user.OutputChannel) {
			discordBot.SendMessage("User is not registered on watch list", m.ChannelID, s)
			return err
		}

		color, err := strconv.ParseInt(args[1], 16, 64)
		if err != nil {
			discordBot.SendMessage("Please provide a valid hex number as color", m.ChannelID, s)
			return err
		}

		err = database.AddOsuWatcherColor(userId, m.ChannelID, int(color))
		if err != nil {
			discordBot.SendMessage("An error occured! Please try again later", m.ChannelID, s)
			return err
		}

		discordBot.SendMessage("Color has been added!", m.ChannelID, s)
		return nil
	},
}
