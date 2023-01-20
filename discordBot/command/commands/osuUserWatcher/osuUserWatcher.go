package osuUserWatcher

import (
	"Digobo/database"
	"Digobo/discordBot/command"
	"github.com/spf13/cobra"
)

var osuUserWatcher = &cobra.Command{
	Use:   "osuuserwatcher",
	Short: "Osu User Watcher",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	osuUserWatcher.AddCommand(addOsuUserWatcher)
	osuUserWatcher.AddCommand(removeOsuUserWatcher)
	osuUserWatcher.AddCommand(colorOsuUserWatcher)
	osuUserWatcher.AddCommand(listOsuUserWatcher)
	command.RootCommand.AddCommand(osuUserWatcher)
}

func hasChannel(channelId string, channelList []database.OsuOutputChannel) bool {
	for _, b := range channelList {
		if b.ChannelId == channelId {
			return true
		}
	}
	return false
}
