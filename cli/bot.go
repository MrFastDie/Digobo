package cli

import (
	"Digobo/discordBot"
	"Digobo/log"
	"github.com/spf13/cobra"

	// load commands when we use the bot
	_ "Digobo/discordBot/command/commands/help"
	_ "Digobo/discordBot/command/commands/ping"
)

var serveCmd = &cobra.Command{
	Use: "start",
	Short: "Run Digobo",
	Long: `run starts the Discord bot service`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info.Println("Starting Discord bot Digobo")

		discordBot.Run()
	},
}


func init() {
	Root.AddCommand(serveCmd)
}