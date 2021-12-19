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
	"github.com/spf13/cobra"
	"time"
)

var addChannel = &cobra.Command{
	Use:   "add [login_name]",
	Short: "adds a user to the watch list",
	Long:  "Add a streamer to get his live notifications",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := command.CommandS
		m := command.CommandM

		user := twitch2.GetUserByLogin(args[0])
		if 0 == len(user.Data) {
			discordBot.SendMessage("No user found with provided login name", m.ChannelID, s)
			return errors.New("no twitch user found")
		}

		_, err := database.GetTwitchWatcherByUserAndChannel(user.Data[0].Id, m.ChannelID)
		if err != sql.ErrNoRows {
			discordBot.SendMessage("An error occurred, please try again later.", m.ChannelID, s)
			return err
		} else if err == nil {
			discordBot.SendMessage("The user has already been added to this channel.", m.ChannelID, s)
			return nil
		}

		err = database.AddTwitchWatcher(user.Data[0].Id, m.ChannelID, false)
		if err != nil {
			discordBot.SendMessage("An error occurred, please try again later.", m.ChannelID, s)
			return err
		}

		jobData := twitchOnline.Data{
			UserId:    user.Data[0].Id,
			ChannelId: m.ChannelID,
		}

		jobDataBytes, _ := json.Marshal(jobData)

		job := scheduler.Job{
			ExecutionTime: time.Now(),
			ExecutionFunc: &twitchOnline.TwitchOnline{},
			Data:          string(jobDataBytes),
		}

		scheduler.GetScheduler().AddScheduledJob(job)

		discordBot.SendMessage(args[0]+" has been added to this channel", m.ChannelID, s)
		return nil
	},
}
