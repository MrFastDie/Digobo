package discordBot

import (
	"Digobo/config"
	"Digobo/discordBot/command"
	"Digobo/log"
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/kballard/go-shellquote"
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

	log.Debug.Printf("Received message: %s\n", m.Content)

	// TODO check if msg of user could be an answer to smth if we have modals (middlewares)
	// TODO specific server settings through middleware such as languages

	if !strings.HasPrefix(m.Content, config.Config.Bot.CommandPrefix) {
		return
	}
	commandStr := strings.TrimPrefix(m.Content, config.Config.Bot.CommandPrefix)
	commandArgs, err := shellquote.Split(commandStr)
	if err != nil {
		err = SendMessage("quote in commend must be closed!", m.ChannelID, s)
		if err != nil {
			log.Error.Println(err)
			return
		}
	}

	newContext := context.Background()
	newContext = context.WithValue(newContext, "s", s)
	newContext = context.WithValue(newContext, "m", m)

	command.RootCommand.SetArgs(commandArgs)
	err = command.RootCommand.ExecuteContext(newContext)
	if err != nil {
		log.Warning.Println("can't execute command", err)
		return
	}
}
