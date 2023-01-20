package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

var CommandI *discordgo.InteractionCreate
var CommandS *discordgo.Session

var RootCommand = &cobra.Command{
	Use:   "digobo",
	Short: "digobo the Discord Go Bot",
	Long:  "A discord bot build in go to serve every need on its own discord server",
}
