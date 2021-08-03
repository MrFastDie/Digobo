package randomGayLinkAnswer

import (
	"Digobo/database"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

var randomGayLinkAnswer = &cobra.Command{
	Use:    "\U0001F970",
	Short:  "Provides a link to a gay porn",
	Long:   "This command provides you a link to a gay porn video",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := cmd.Context().Value("s").(*discordgo.Session)
		m := cmd.Context().Value("m").(*discordgo.MessageCreate)

		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return nil
		}

		link, err := database.GetKeyValuePairByCommand(cmd.Use)
		if err != nil {
			log.Error.Println("cant fetch a random gay link from DB", err)
			return err
		}

		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Error.Println("can't fetch channel", err)
			return err
		}

		if !channel.NSFW {
			return nil
		}

		_, err = s.ChannelMessageSend(m.ChannelID, link)
		if err != nil {
			log.Error.Println("can't send embed", err)
			return err
		}

		return nil
	},
}

func init() {
	randomGayLinkAnswer.AddCommand(addRandomGayLinkAnswer)
	randomGayLinkAnswer.AddCommand(removeRandomGayLinkAnswer)
	command.RootCommand.AddCommand(randomGayLinkAnswer)
}
