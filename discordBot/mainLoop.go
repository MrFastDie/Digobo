package discordBot

import (
	"Digobo/config"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

func mainLoop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		execCommands(s, i)
		break
	case discordgo.InteractionMessageComponent:
		// To get the custom ID of that message
		//SendInteractionMessage(i.MessageComponentData().CustomID, s, i)
		SendInteractionMessage("Interaction type not implemented", s, i)
		break
	default:
		SendInteractionMessage("Interaction type not implemented", s, i)
	}
}

func execCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandStr := []string{i.ApplicationCommandData().Name}

	applicationName := i.ApplicationCommandData().Name
	if "" != config.Config.Bot.CommandPrefix {
		reStr := regexp.MustCompile("^(.*?)" + config.Config.Bot.CommandPrefix + "-(.*)$")
		repStr := "${1}$2"
		applicationName = reStr.ReplaceAllString(applicationName, repStr)
	}

	if len(i.ApplicationCommandData().Options) > 0 {
		commandStr = append(commandStr, i.ApplicationCommandData().Options[0].Name)
		commandStr = append(commandStr, strings.Split(i.ApplicationCommandData().Options[0].StringValue(), " ")...)
	}

	log.Debug.Printf("Received slash command: %s\n", commandStr)

	rootCmd, ok := command.GetMap()[applicationName]
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
