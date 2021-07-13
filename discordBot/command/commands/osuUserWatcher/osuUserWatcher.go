package osuUserWatcher

import (
	"Digobo/apps/osu"
	"Digobo/database"
	"Digobo/discordBot/command"
	"Digobo/log"
	CrawlOsuProfiles "Digobo/scheduler/jobs/crawlOsuProfiles"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"time"
)

type OsuUserWatcher struct{}

func (this *OsuUserWatcher) Name() string {
	return "OsuUserWatcher"
}

func (this *OsuUserWatcher) Title() string {
	return "Adds or remove a user from the watch list"
}

func (this *OsuUserWatcher) Description() string {
	return "OsuUserWatcher add/remove <osu_user_id> (adds or removes a user to the watch list of this channel)"
}

func (this *OsuUserWatcher) Hidden() bool {
	return false
}

func (this *OsuUserWatcher) HasInteractions() bool {
	return false
}

func (this *OsuUserWatcher) Execute(args string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return nil
	}

	parts := strings.SplitN(args, " ", 2)

	if parts[0] == "add" {
		userAlreadyPresent := true
		userId, err := strconv.Atoi(parts[1])
		if err != nil {
			sendChannelMessage("Please provide a valid user_id", s, m)
		}

		user, err := database.GetOsuWatcher(userId)
		if err != nil && err.Error() == database.NO_ROWS {
			osuUser, err := osu.GetUser(userId)
			if err != nil {
				log.Info.Println("can't fetch osu! user from API")
				sendChannelMessage("No user found with provided user_id", s, m)
				return err
			}

			err = database.AddOsuWatcherUser(userId, osuUser.Username)
			if err != nil {
				log.Error.Println("can't add osu! user to DB", err)
				sendChannelMessage("An error occurred, please try again later", s, m)
				return err
			}

			userAlreadyPresent = false
			user, _ = database.GetOsuWatcher(userId)
		} else if err != nil {
			log.Error.Println("can't fetch osu! user from db", err)
			sendChannelMessage("An error occurred, please try again later", s, m)
			return err
		}

		err = database.AddOsuWatcherOutputChannel(userId, m.ChannelID)
		if err != nil && strings.Contains(err.Error(), database.PQ_DUPLICATES) {
			sendChannelMessage("User has already been added", s, m)
			return nil
		} else if err != nil {
			log.Error.Println("can't add channel to osu! user in db", err)
			sendChannelMessage("An error occurred, please try again later", s, m)
			return err
		}

		if !userAlreadyPresent {
			CrawlerData := CrawlOsuProfiles.Data{
				UserId:        userId,
				UserName:      user.UserName,
				OutputChannel: []string{m.ChannelID},
			}
			CrawlStrData, _ := json.Marshal(CrawlerData)
			CrawlOsuProfiles.CrawlOsuProfilesJobStart(time.Now(), string(CrawlStrData))
		}

		sendChannelMessage(user.UserName+" has been added to this channel", s, m)
	}

	if parts[0] == "remove" {
		userId, err := strconv.Atoi(parts[1])
		if err != nil {
			sendChannelMessage("Please provide a valid user_id", s, m)
		}

		err = database.RemoveOsuWatcherOutputChannel(userId, m.ChannelID)
		if err != nil {
			log.Error.Println("cant delete output channel from DB", err)
			sendChannelMessage("An error occurred, please try again later", s, m)
			return err
		}

		sendChannelMessage("Removal has been successful", s, m)
	}

	return nil
}

func sendChannelMessage(msg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		log.Error.Println("can't send channel message", err)
	}
}

func init() {
	command.Commands.LoadCommand(&OsuUserWatcher{})
}
