package ping

import (
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"github.com/kr/pretty"
	"time"
)

type Ping struct{}

func (this *Ping) Name() string {
	return "ping"
}

func (this *Ping) Title() string {
	return "Ping! example"
}

func (this *Ping) Description() string {
	return "This is an example module for commands"
}

func (this *Ping) HasInteractions() bool {
	return false
}

func (this *Ping) Execute(args string, s *discordgo.Session, m *discordgo.MessageCreate) error {
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
		Color:       0x00ff00,
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

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		message, _ := s.ChannelMessageSendEmbed(answerChannelID, embed)
		pretty.Println(m)
		pretty.Println(message)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(answerChannelID, "Ping!")
	}

	return nil
}

func init() {
	command.Commands.LoadCommand(&Ping{})
}
