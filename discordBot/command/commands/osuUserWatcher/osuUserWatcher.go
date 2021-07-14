package osuUserWatcher

import (
	"Digobo/apps/osu"
	"Digobo/database"
	"Digobo/discordBot/command"
	"Digobo/log"
	CrawlOsuProfiles "Digobo/scheduler/jobs/crawlOsuProfiles"
	"encoding/json"
	"fmt"
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
	return "add <osu_user_id> - Adds a user to the watch list\n" +
		"remote <osu_user_id> - Removes a user from the watch list\n" +
		"color <osu_user_id> <color_hex_code> - Changes the color of the message embed\n" +
		"list - List all members of the watch list"
}

func (this *OsuUserWatcher) Description() string {
	return "add <osu_user_id> - Adds a user to the watch list\n" +
		"remote <osu_user_id> - Removes a user from the watch list\n" +
		"color <osu_user_id> <color_hex_code> - Changes the color of the message embed\n" +
		"list - List all members of the watch list"
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

	switch parts[0] {
	case "add":
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
				UserId:   userId,
				UserName: user.UserName,
				OutputChannel: []database.OsuOutputChannel{
					{
						ChannelId: m.ChannelID,
					},
				},
			}
			CrawlStrData, _ := json.Marshal(CrawlerData)
			CrawlOsuProfiles.CrawlOsuProfilesJobStart(time.Now(), string(CrawlStrData))
		}

		sendChannelMessage(user.UserName+" has been added to this channel", s, m)

		break

	case "remove":
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

		break

	case "color":
		colorParts := strings.SplitN(parts[1], " ", 2)

		userId, err := strconv.Atoi(colorParts[0])
		if err != nil {
			sendChannelMessage("Please provide a valid user_id", s, m)
		}

		user, err := database.GetOsuWatcher(userId)
		if err != nil || !hasChannel(m.ChannelID, user.OutputChannel) {
			sendChannelMessage("User is not registered on watch list", s, m)
			return err
		}

		color, err := strconv.ParseInt(colorParts[1], 16, 64)
		if err != nil {
			sendChannelMessage("Please provide a valid hex number as color", s, m)
			return err
		}

		err = database.AddOsuWatcherColor(userId, m.ChannelID, int(color))
		if err != nil {
			sendChannelMessage("An error occured! Please try again later", s, m)
			return err
		}

		sendChannelMessage("Color has been added!", s, m)

		break

	case "list":
		list, err := database.GetOsuWatcherListByChannel(m.ChannelID)
		if err != nil && err.Error() == database.NO_ROWS {
			sendChannelMessage("There are currently no member in your watch list", s, m)
			return nil
		} else if err != nil {
			log.Error.Println("can't fetch osu watcher list", err)
			return err
		}

		for _, entry := range list {
			sendChannelMessage(fmt.Sprintf("%s with ID: %d (%s/users/%d)", entry.UserName, entry.UserId, osu.OSU_API_URL, entry.UserId), s, m)
		}

		break
	}

	return nil
}

func sendChannelMessage(msg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		log.Error.Println("can't send channel message", err)
	}
}

func hasChannel(channelId string, channelList []database.OsuOutputChannel) bool {
	for _, b := range channelList {
		if b.ChannelId == channelId {
			return true
		}
	}
	return false
}

func init() {
	command.Commands.LoadCommand(&OsuUserWatcher{})
}
