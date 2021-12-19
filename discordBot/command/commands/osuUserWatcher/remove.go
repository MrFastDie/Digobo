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
	Use:   "remove [user_id]",
	Short: "removes a user from the watch list",
	Long:  "This command allows you to remove a user from your personal watch list by a given id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		m := command.CommandM

		userId, err := strconv.Atoi(args[0])
		if err != nil {
			discordBot.SendMessage("Please provide a valid user_id", m.ChannelID, s)
			return err
		}

		err = database.RemoveOsuWatcherOutputChannel(userId, m.ChannelID)
		if err != nil {
			log.Error.Println("cant delete output channel from DB", err)
			discordBot.SendMessage("An error occurred, please try again later", m.ChannelID, s)
			return err
		}

		discordBot.SendMessage("Removal has been successful", m.ChannelID, s)
		return nil
	},
}
