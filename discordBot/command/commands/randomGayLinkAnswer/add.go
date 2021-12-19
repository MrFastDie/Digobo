package randomGayLinkAnswer

import (
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/spf13/cobra"
)

var addRandomGayLinkAnswer = &cobra.Command{
	Use:    "add [url]",
	Short:  "Adds a new gay porn link",
	Long:   "This command adds a new gay porn link",
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

		err := database.AddKeyValuePairByCommand(cmd.Parent().Use, args[0], m.Author.ID)
		if err != nil {
			log.Error.Println("cant add random gay link to DB", err)
			return err
		}

		discordBot.SendMessage(fmt.Sprintf("Link %s successfully added", args[0]), m.ChannelID, s)
		return nil
	},
}

func init() {
	command.RootCommand.AddCommand(randomGayLinkAnswer)
}
