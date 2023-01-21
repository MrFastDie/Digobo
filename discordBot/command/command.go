package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

type Command struct {
	Name        string
	Description string

	SubCommands map[string]SubCommand

	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

var RootCommand = &cobra.Command{
	Use:   "digobo",
	Short: "digobo the Discord Go Bot",
	Long:  "A discord bot build in go to serve every need on its own discord server",
}
