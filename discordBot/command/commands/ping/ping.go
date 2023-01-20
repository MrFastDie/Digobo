package ping

import (
	"Digobo/config"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
	"time"
)

var pingCommand = &cobra.Command{
	Use:   "ping",
	Short: "Ping! example",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		i := command.CommandI

		embeds := []*discordgo.MessageEmbed{{
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
		}}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: embeds,
			},
		})
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
