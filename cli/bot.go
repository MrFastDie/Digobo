package cli

import (
	"Digobo/discordBot"
	"Digobo/log"
	CrawlOsuProfiles "Digobo/scheduler/jobs/crawlOsuProfiles"
	"Digobo/scheduler/jobs/crawlReminderEvents"
	"encoding/json"
	"github.com/spf13/cobra"
	"time"

	// load commands when we use the bot
	_ "Digobo/discordBot/command/commands/help"
	_ "Digobo/discordBot/command/commands/ping"
	_ "Digobo/discordBot/command/commands/randomGayLinkAnswer"

	_ "Digobo/scheduler"
)

var serveCmd = &cobra.Command{
	Use:   "start",
	Short: "Run Digobo",
	Long:  `run starts the Discord bot service`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info.Println("Starting Discord bot Digobo")

		crawlReminderEvents.CrawlReminderEventJobStart(time.Now(), "")

		const TEST_CHANNEL = "835175407216623727"
		const BROADCAST_CHANNEL = "863777071113043989"

		createCrawlEvent(
			15817570,
			"MrFastDie",
			BROADCAST_CHANNEL,
		)

		createCrawlEvent(
			22106410,
			"AmateurUwU",
			BROADCAST_CHANNEL,
		)

		createCrawlEvent(
			13467065,
			"-Darius",
			BROADCAST_CHANNEL,
		)

		createCrawlEvent(
			19490974,
			"the50sten",
			BROADCAST_CHANNEL,
		)

		discordBot.Run()
	},
}

func createCrawlEvent(userId int, userName string, outputChannel string) {
	CrawlerData := CrawlOsuProfiles.Data{
		UserId:        userId,
		UserName:      userName,
		OutputChannel: outputChannel,
	}
	CrawlStrData, _ := json.Marshal(CrawlerData)
	CrawlOsuProfiles.CrawlOsuProfilesJobStart(time.Now(), string(CrawlStrData))
}

func init() {
	Root.AddCommand(serveCmd)
}
