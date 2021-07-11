package cli

import (
	"Digobo/database"
	"Digobo/log"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate DB",
	Long:  `migrates the DB`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info.Println("Migrating DB")

		database.Migrate()
	},
}

func init() {
	Root.AddCommand(migrateCmd)
}
