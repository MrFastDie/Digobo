package ping

import (
	"Digobo/config"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
	"time"
)

var pingCommand = &cobra.Command{
	Use:   "ping",
	Short: "Ping! example",
	Long:  "This is an example module for commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := cmd.Context().Value("s").(*discordgo.Session)
		m := cmd.Context().Value("m").(*discordgo.MessageCreate)

		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return nil
		}

		answerChannelID := m.ChannelID
		isDm, err := discordBot.ComesFromDM(s, m)
		if err != nil {
			log.Error.Println("cant check if channel is DM", err)
			return err
		}

		if !isDm {
			channel, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error.Println("cant create a channel from author ID", err)
				return err
			}

			answerChannelID = channel.ID
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Pong!",
			Description: "This is the Pong! embed",
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       config.Config.Bot.DefaultEmbedColor,
			Author:      &discordgo.MessageEmbedAuthor{},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Field 1",
					Value:  "Field 1",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Field 2",
					Value:  "Field 2",
					Inline: true,
				},
			},
		}

		_, err = s.ChannelMessageSendEmbed(answerChannelID, embed)
		if err != nil {
			log.Error.Println("can't send embed", err)
			return err
		}

		return nil
	},
}

func init() {
	command.RootCommand.AddCommand(pingCommand)
}
