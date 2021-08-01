package command

import (
	"github.com/spf13/cobra"
)

// TODO implement event system with states
type Command interface {
	// Does this command use Events or is it just an ping - pong command?
	HasInteractions() bool

	GetCobraCommand() *cobra.Command
}

var RootCommand = &cobra.Command{
	Use:   "digobo",
	Short: "digobo the Discord Go Bot",
	Long:  "A discord bot build in go to serve every need on its own discord server",
}

func LoadCommand(command Command) {
	RootCommand.AddCommand(command.GetCobraCommand())
}

func init() {

}
