package command

import "github.com/bwmarrin/discordgo"

type SubCommand struct {
	Name        string
	Description string

	Type discordgo.ApplicationCommandOptionType

	SubCommands map[string]SubCommand

	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate, args interface{}) error
}
