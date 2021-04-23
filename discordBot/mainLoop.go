package discordBot

import (
	"Digobo/config"
	"Digobo/discordBot/command"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func mainLoop(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// TODO check if msg of user could be an answer to smth if we have modals (middlewares)
	// TODO specific server settings through middleware such as languages

	commandParts := strings.SplitN(m.Content, " ", 2)
	if !strings.HasPrefix(commandParts[0], config.Config.Bot.CommandPrefix) {
		return
	}

	commandParts[0] = strings.TrimPrefix(commandParts[0], config.Config.Bot.CommandPrefix)

	command, err := command.Commands.GetCommand(commandParts[0])
	if err != nil {
		return
	}

	var args string
	if len(commandParts) > 1 {
		args = commandParts[1]
	}

	err = command.Execute(args, s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "An error occured - please try again")
	}
}
