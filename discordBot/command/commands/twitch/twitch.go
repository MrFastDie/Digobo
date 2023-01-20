package twitch

import (
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"github.com/spf13/cobra"
)

var twitch = &cobra.Command{
	Use:   "twitch",
	Short: "Twitch actions",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		i := command.CommandI

		return discordBot.SendInteractionMessage("Basic twitch command executed", s, i)
	},
}

func init() {
	twitch.AddCommand(addChannel)
	command.RootCommand.AddCommand(twitch)
}
