package randomGayLinkAnswer

import (
	"Digobo/database"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type RandomGayLinkAnswer struct{}

func (this *RandomGayLinkAnswer) Name() string {
	return "\U0001F970"
}

func (this *RandomGayLinkAnswer) Title() string {
	return "Random gay porn link"
}

func (this *RandomGayLinkAnswer) Description() string {
	return "Returns a random gay porn link (include a + or - with a following link to add or delete it in the database)"
}

func (this *RandomGayLinkAnswer) Hidden() bool {
	return true
}

func (this *RandomGayLinkAnswer) HasInteractions() bool {
	return false
}

func (this *RandomGayLinkAnswer) Execute(args string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return nil
	}

	link, err := database.GetRandomAnswerListValueByCommand(this.Name())
	if err != nil {
		log.Error.Println("cant fetch a random gay link from DB", err)
		return err
	}

	if args != "" && command.IsBotMaster(m.Author.ID) {
		parts := strings.SplitN(args, " ", 2)

		if parts[0] == "+" {
			err := database.AddRandomAnswerListValueByCommand(this.Name(), parts[1], m.Author.ID)
			if err != nil {
				log.Error.Println("cant add random gay link to DB", err)
				return err
			}

			link = fmt.Sprintf("Link %s successfully added", parts[1])
		}

		if parts[0] == "-" {
			err := database.RemoveRandomAnswerListValueByCommandAndValue(this.Name(), parts[1])
			if err != nil {
				log.Error.Println("cant delete random gay link from DB", err)
				return err
			}

			link = fmt.Sprintf("Link %s successfully deleted", parts[1])
		}
	}

	s.ChannelMessageSend(m.ChannelID, link)

	return nil
}

func init() {
	command.Commands.LoadCommand(&RandomGayLinkAnswer{})
}
