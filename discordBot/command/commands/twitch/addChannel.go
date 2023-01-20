package twitch

import (
	twitch2 "Digobo/apps/twitch"
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/scheduler"
	"Digobo/scheduler/jobs/twitchOnline"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var addChannel = &cobra.Command{
	Use:   "add",
	Short: "adds a user to the watch list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		i := command.CommandI

		discordBot.SendInteractionMessage("Command received", s, i)

		user := twitch2.GetUserByLogin(args[0])
		if 0 == len(user.Data) {
			discordBot.SendMessageFromInteraction("No user found with provided login name", s, i)
			return errors.New("no twitch user found")
		}

		_, err := database.GetTwitchWatcherByUserAndChannel(user.Data[0].Id, i.ChannelID)
		if err != sql.ErrNoRows {
			discordBot.SendMessageFromInteraction("An error occurred, please try again later.", s, i)
			return err
		} else if err == nil {
			discordBot.SendMessageFromInteraction("The user has already been added to this channel.", s, i)
			return nil
		}

		err = database.AddTwitchWatcher(user.Data[0].Id, i.ChannelID, false)
		if err != nil {
			discordBot.SendMessageFromInteraction("An error occurred, please try again later.", s, i)
			return err
		}

		jobData := twitchOnline.Data{
			UserId:    user.Data[0].Id,
			ChannelId: i.ChannelID,
		}

		jobDataBytes, _ := json.Marshal(jobData)

		job := scheduler.Job{
			ExecutionTime: time.Now(),
			ExecutionFunc: &twitchOnline.TwitchOnline{},
			Data:          string(jobDataBytes),
		}

		scheduler.GetScheduler().AddScheduledJob(job)

		return discordBot.SendMessageFromInteraction(fmt.Sprintf("%s has been added to this channel", args[0]), s, i)
	},
}
