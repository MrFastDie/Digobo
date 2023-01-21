package osuUserWatcher

import (
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

var Remove = command.SubCommand{
	Name:        "remove",
	Description: "Removes a user from the watch list",
	Type:        discordgo.ApplicationCommandOptionString,
	SubCommands: nil,
	Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate, args interface{}) error {
		discordBot.SendInteractionMessage("Command received", s, i)

		userId, err := strconv.Atoi(args.(string))
		if err != nil {
			discordBot.SendMessage("Please provide a valid user_id", i.ChannelID, s)
			return err
		}

		err = database.RemoveOsuWatcherOutputChannel(userId, i.ChannelID)
		if err != nil {
			log.Error.Println("cant delete output channel from DB", err)
			discordBot.SendMessage("An error occurred, please try again later", i.ChannelID, s)
			return err
		}

		discordBot.SendMessage("Removal has been successful", i.ChannelID, s)
		return nil
	},
}
