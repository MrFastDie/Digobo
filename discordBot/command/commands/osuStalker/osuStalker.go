package ping

import (
	"Digobo/config"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"time"
)

type OsuStalker struct{}

func (this *OsuStalker) Name() string {
	return "OsuStalker"
}

func (this *OsuStalker) Title() string {
	return "Add a player to stalk when he has uploaded new maps"
}

func (this *OsuStalker) Description() string {
	return "Stalks a player and see if he has uploaded a new map"
}

func (this *OsuStalker) Hidden() bool {
	return false
}

func (this *OsuStalker) HasInteractions() bool {
	return false
}

func (this *OsuStalker) Execute(args string, s *discordgo.Session, m *discordgo.MessageCreate) error {
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
		Fields:      []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Field 1",
				Value: "Field 1",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name: "Field 2",
				Value: "Field 2",
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
}

func init() {
	command.Commands.LoadCommand(&OsuStalker{})
}
