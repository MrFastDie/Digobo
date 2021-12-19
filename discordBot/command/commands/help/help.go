package help

import (
	"Digobo/config"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var helpCommand = &cobra.Command{
	Use:   "help",
	Short: "Help about any command",
	Long: `Help provides help for any command in the bot.
Simply type ` + config.Config.Bot.CommandPrefix + `help [path to command] for full details.`,
}

func init() {
	command.RootCommand.SetHelpCommand(helpCommand)
	command.RootCommand.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		s := command.CommandS
		m := command.CommandM

		args = args[1:]

		commands := command.RootCommand.Commands()
		usage := config.Config.Bot.CommandPrefix + "[command]"
		if len(args) > 0 {
			usage = config.Config.Bot.CommandPrefix + strings.Join(args, " ") + " [command]"
			cmd = command.RootCommand
			for _, arg := range args {
				found := false
				for _, subCmd := range cmd.Commands() {
					if subCmd.Use == arg {
						cmd = subCmd
						found = true
						break
					}
				}

				if !found {
					err := discordBot.SendMessage("Command "+config.Config.Bot.CommandPrefix+strings.Join(args, " ")+" not found!", m.ChannelID, s)
					if err != nil {
						log.Error.Println("can't send message", err)
						return
					}
				}
			}

			commands = cmd.Commands()
		}

		var fields []*discordgo.MessageEmbedField

		var commandFields []*discordgo.MessageEmbedField
		for _, cmd := range commands {
			if cmd.Hidden == true {
				continue
			}

			commandFields = append(commandFields, &discordgo.MessageEmbedField{
				Name:   discordBot.Whitespace + cmd.Use,
				Value:  discordBot.Whitespace + cmd.Short,
				Inline: false,
			})
		}

		if len(commandFields) > 0 {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   "Available Commands",
				Value:  "Below is a list of available commands",
				Inline: false,
			})

			fields = append(fields, commandFields...)
		}

		if cmd.Flags().FlagUsages() != "" {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   "Available Flags",
				Value:  cmd.Flags().FlagUsages(),
				Inline: false,
			})
		}

		embed := &discordgo.MessageEmbed{
			Title:       cmd.Use,
			Description: cmd.Long + "\n\nUsage:\n" + discordBot.Whitespace + usage,
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       config.Config.Bot.DefaultEmbedColor,
			Author:      &discordgo.MessageEmbedAuthor{},
			Fields:      fields,
		}

		err := discordBot.SendEmbed(embed, m.ChannelID, s)
		if err != nil {
			log.Error.Println("can't send embed", err)
		}
	})
}
