package discordBot

import (
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func mainLoop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandStr := []string{i.ApplicationCommandData().Name}

	if len(i.ApplicationCommandData().Options) > 0 {
		commandStr = append(commandStr, i.ApplicationCommandData().Options[0].Name)
		commandStr = append(commandStr, strings.Split(i.ApplicationCommandData().Options[0].StringValue(), " ")...)
	}

	log.Debug.Printf("Received slash command: %s\n", commandStr)

	command.CommandI = i
	command.CommandS = s

	command.RootCommand.SetArgs(commandStr)
	err := command.RootCommand.Execute()
	if err != nil {
		log.Warning.Println("can't execute command", err)
		return
	}
}
