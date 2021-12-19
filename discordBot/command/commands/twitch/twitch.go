package twitch

import (
	"Digobo/discordBot/command"
	"github.com/spf13/cobra"
)

var twitch = &cobra.Command{
	Use:   "twitch",
	Short: "Twitch actions",
	Long:  "Execute specific twitch actions",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	twitch.AddCommand(addChannel)
	command.RootCommand.AddCommand(twitch)
}
