package help

import (
	"Digobo/discordBot/command"
	"github.com/bwmarrin/discordgo"
	"time"
)

type Help struct{}

func (this *Help) Name() string {
	return "help"
}

func (this *Help) Title() string {
	return "Show this message"
}

func (this *Help) Description() string {
	return "This command is ment to be helping you finding the right commands"
}

func (this *Help) HasInteractions() bool {
	return false
}

func (this *Help) Execute(args string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return nil
	}

	answerChannelID := m.ChannelID

	var fields []*discordgo.MessageEmbedField

	for k, v := range command.Commands.GetCommands() {
		fields = append(fields, &discordgo.MessageEmbedField{Name: k, Value: v.Title(), Inline: false})
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Help",
		Description: this.Description(),
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0x00ff00,
		Author:      &discordgo.MessageEmbedAuthor{},
		Fields:      fields,
	}

	s.ChannelMessageSendEmbed(answerChannelID, embed)

	return nil
}

func init() {
	command.Commands.LoadCommand(&Help{})
}
