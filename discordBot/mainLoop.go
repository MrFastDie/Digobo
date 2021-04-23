package discordBot

import (
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

	commandParts := strings.SplitN(m.Content, " ", 1)
	command, err := command.Commands.GetCommand(commandParts[0])
	if err != nil {
		return
	}

	var args string
	if len(commandParts) > 1 {
		args = commandParts[1]
	}

	command.Execute(args, s, m)
}
