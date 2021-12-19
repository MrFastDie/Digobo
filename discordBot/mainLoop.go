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
		err = SendMessage("quote in command must be closed!", m.ChannelID, s)
		if err != nil {
			log.Error.Println(err)
			return
		}
	}

	ctx := command.RootCommand.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	command.CommandM = m
	command.CommandS = s
	//ctx = context.WithValue(ctx, "s", s)
	//ctx = context.WithValue(ctx, "m", m)
	//ctx, ctxCancel := context.WithCancel(ctx)
	//defer ctxCancel()

	command.RootCommand.SetArgs(commandArgs)
	err = command.RootCommand.Execute()
	//err = command.RootCommand.ExecuteContext(ctx)
	if err != nil {
		log.Warning.Println("can't execute command", err)
		return
	}
}
