package randomGayLinkAnswer

import (
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/spf13/cobra"
)

var removeRandomGayLinkAnswer = &cobra.Command{
	Use:    "remove [url]",
	Short:  "Remove a gay porn vid based on its link",
	Long:   "This command removes a gay porn link",
	Hidden: true,
	Args:   cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		m := command.CommandM

		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return nil
		}

		if !command.IsBotMaster(m.Author.ID) {
			return nil
		}

		err := database.RemoveKeyValuePairByCommandAndValue(cmd.Parent().Use, args[0])
		if err != nil {
			log.Error.Println("cant delete random gay link from DB", err)
			return err
		}

		discordBot.SendMessage(fmt.Sprintf("Link %s successfully deleted", args[0]), m.ChannelID, s)

		return nil
	},
}

func init() {
	command.RootCommand.AddCommand(randomGayLinkAnswer)
}
