package discordBot

import (
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
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

	rootCmd, ok := command.Map[i.ApplicationCommandData().Name]
	if !ok {
		SendInteractionMessage(fmt.Sprintf("Unknown command %s", i.ApplicationCommandData().Name), s, i)
		return
	}

	if len(i.ApplicationCommandData().Options) > 0 {
		for _, subCommand := range i.ApplicationCommandData().Options {
			cmd, ok := rootCmd.SubCommands[subCommand.Name]
			if !ok {
				SendMessage(fmt.Sprintf("Unknown subcommand %s", subCommand.Name), i.ChannelID, s)
				continue
			}

			cmd.Execute(s, i, subCommand.Value)
		}

		return
	}

	rootCmd.Execute(s, i)
}
