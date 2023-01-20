package osuUserWatcher

import (
	"Digobo/apps/osu"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	CrawlOsuProfiles "Digobo/scheduler/jobs/crawlOsuProfiles"
	"encoding/json"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

var addOsuUserWatcher = &cobra.Command{
	Use:   "add",
	Short: "adds a user to the watch list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		i := command.CommandI

		discordBot.SendInteractionMessage("Command received", s, i)

		userAlreadyPresent := true
		userId, err := strconv.Atoi(args[0])
		if err != nil {
			discordBot.SendMessage("Please provide a valid user_id", i.ChannelID, s)
			return err
		}

		user, err := database.GetOsuWatcher(userId)
		if err != nil && err.Error() == database.NO_ROWS {
			osuUser, err := osu.GetUser(userId)
			if err != nil {
				log.Info.Println("can't fetch osu! user from API")
				discordBot.SendMessage("No user found with provided user_id", i.ChannelID, s)
				return err
			}

			err = database.AddOsuWatcherUser(userId, osuUser.Username)
			if err != nil {
				log.Error.Println("can't add osu! user to DB", err)
				discordBot.SendMessage("An error occurred, please try again later", i.ChannelID, s)
				return err
			}

			userAlreadyPresent = false
			user, _ = database.GetOsuWatcher(userId)
		} else if err != nil {
			log.Error.Println("can't fetch osu! user from db", err)
			discordBot.SendMessage("An error occurred, please try again later", i.ChannelID, s)
			return err
		}

		err = database.AddOsuWatcherOutputChannel(userId, i.ChannelID)
		if err != nil && strings.Contains(err.Error(), database.PQ_DUPLICATES) {
			discordBot.SendMessage("User has already been added", i.ChannelID, s)
			return nil
		} else if err != nil {
			log.Error.Println("can't add channel to osu! user in db", err)
			discordBot.SendMessage("An error occurred, please try again later", i.ChannelID, s)
			return err
		}

		if !userAlreadyPresent {
			CrawlerData := CrawlOsuProfiles.Data{
				UserId:   userId,
				UserName: user.UserName,
				Retries:  0,
				OutputChannel: []database.OsuOutputChannel{
					{
						ChannelId: i.ChannelID,
					},
				},
			}
			CrawlStrData, _ := json.Marshal(CrawlerData)
			CrawlOsuProfiles.CrawlOsuProfilesJobStart(time.Now(), string(CrawlStrData))
		}

		discordBot.SendMessage(user.UserName+" has been added to this channel", i.ChannelID, s)
		return nil
	},
}
