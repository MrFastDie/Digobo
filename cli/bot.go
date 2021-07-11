package cli

import (
	"Digobo/discordBot"
	"Digobo/log"
	CrawlOsuProfiles "Digobo/scheduler/jobs/crawlOsuProfiles"
	"Digobo/scheduler/jobs/crawlReminderEvents"
	"github.com/spf13/cobra"
	"time"

	// load commands when we use the bot
	_ "Digobo/discordBot/command/commands/help"
	_ "Digobo/discordBot/command/commands/osuStalker"
	_ "Digobo/discordBot/command/commands/ping"
	_ "Digobo/discordBot/command/commands/randomGayLinkAnswer"

	_ "Digobo/scheduler"
)

var serveCmd = &cobra.Command{
	Use: "start",
	Short: "Run Digobo",
	Long: `run starts the Discord bot service`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info.Println("Starting Discord bot Digobo")

		crawlReminderEvents.CrawlReminderEventJobStart(time.Now(), "")
		CrawlOsuProfiles.CrawlOsuProfilesJobStart(time.Now(), "")

		discordBot.Run()
	},
}

func init() {
	Root.AddCommand(serveCmd)
}