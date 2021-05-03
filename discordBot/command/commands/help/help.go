package help

import (
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
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
	return "The purpose of this command is to help you find the right commands"
}

func (this *Help) Hidden() bool {
	return false
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

	var embed discordgo.MessageEmbed
	var fields []*discordgo.MessageEmbedField

	if args == "" {
		for k, v := range command.Commands.GetCommands() {
			if v.Hidden() {
				continue
			}

			fields = append(fields, &discordgo.MessageEmbedField{Name: k, Value: v.Title(), Inline: false})
		}

		embed = discordgo.MessageEmbed{
			Title:       "Help",
			Description: this.Description(),
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       0x00ff00,
			Author:      &discordgo.MessageEmbedAuthor{},
			Fields:      fields,
		}
	} else {
		getCommand, err := command.Commands.GetCommand(args)
		if err != nil {
			embed = discordgo.MessageEmbed{
				Title:       "Not found",
				Description: fmt.Sprintf("The command %s could not be found!", args),
				Timestamp:   time.Now().Format(time.RFC3339),
				Color:       0x00ff00,
				Author:      &discordgo.MessageEmbedAuthor{},
			}
		} else {
			embed = discordgo.MessageEmbed{
				Title:       getCommand.Title(),
				Description: getCommand.Description(),
				Timestamp:   time.Now().Format(time.RFC3339),
				Color:       0x00ff00,
				Author:      &discordgo.MessageEmbedAuthor{},
			}
		}
	}

	_, err := s.ChannelMessageSendEmbed(answerChannelID, &embed)
	if err != nil {
		log.Error.Println("can't send embed", err)
		return err
	}


	return nil
}

func init() {
	command.Commands.LoadCommand(&Help{})
}
