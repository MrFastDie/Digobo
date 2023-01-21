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
	"github.com/bwmarrin/discordgo"
	"time"
)

var AddChannel = command.SubCommand{
	Name:        "add",
	Description: "adds a given user to the watch list",
	Type:        discordgo.ApplicationCommandOptionString,
	SubCommands: nil,
	Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate, args interface{}) error {
		discordBot.SendInteractionMessage("Command received", s, i)

		user := twitch2.GetUserByLogin(args.(string))
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

		return discordBot.SendMessageFromInteraction(fmt.Sprintf("%s has been added to this channel", args.(string)), s, i)
	},
}
