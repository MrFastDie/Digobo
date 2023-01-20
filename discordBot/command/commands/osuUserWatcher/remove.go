package osuUserWatcher

import (
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/spf13/cobra"
	"strconv"
)

var removeOsuUserWatcher = &cobra.Command{
	Use:   "remove",
	Short: "removes a user from the watch list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		i := command.CommandI

		discordBot.SendInteractionMessage("Command received", s, i)

		userId, err := strconv.Atoi(args[0])
		if err != nil {
			discordBot.SendMessage("Please provide a valid user_id", i.ChannelID, s)
			return err
		}

		err = database.RemoveOsuWatcherOutputChannel(userId, i.ChannelID)
		if err != nil {
			log.Error.Println("cant delete output channel from DB", err)
			discordBot.SendMessage("An error occurred, please try again later", i.ChannelID, s)
			return err
		}

		discordBot.SendMessage("Removal has been successful", i.ChannelID, s)
		return nil
	},
}
