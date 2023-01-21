package osuUserWatcher

import (
	"Digobo/database"
	"Digobo/discordBot"
	"Digobo/discordBot/command"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

var Color = command.SubCommand{
	Name:        "color",
	Description: "Changes the embed color of the given user id: color <userId> <colorHex>",
	Type:        discordgo.ApplicationCommandOptionString,
	SubCommands: nil,
	Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate, oArgs interface{}) error {
		discordBot.SendInteractionMessage("Command received", s, i)

		args := strings.Split(oArgs.(string), " ")

		userId, err := strconv.Atoi(args[0])
		if err != nil {
			discordBot.SendMessage("Please provide a valid user_id", i.ChannelID, s)
		}

		user, err := database.GetOsuWatcher(userId)
		if err != nil || !hasChannel(i.ChannelID, user.OutputChannel) {
			discordBot.SendMessage("User is not registered on watch list", i.ChannelID, s)
			return err
		}

		color, err := strconv.ParseInt(args[1], 16, 64)
		if err != nil {
			discordBot.SendMessage("Please provide a valid hex number as color", i.ChannelID, s)
			return err
		}

		err = database.AddOsuWatcherColor(userId, i.ChannelID, int(color))
		if err != nil {
			discordBot.SendMessage("An error occured! Please try again later", i.ChannelID, s)
			return err
		}

		discordBot.SendMessage("Color has been added!", i.ChannelID, s)
		return nil
	},
}
