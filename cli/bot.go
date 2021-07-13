package cli

import (
	"Digobo/database"
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

		osuWatcher, err := database.GetOsuWatcher()
		if err != nil {
			log.Error.Fatal("can't fetch osu watcher", err)
			return
		}

		for _, watcher := range osuWatcher {
			for _, channel := range watcher.OutputChannel {
				createCrawlEvent(
					watcher.UserId,
					watcher.UserName,
					channel,
				)
			}
		}

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
