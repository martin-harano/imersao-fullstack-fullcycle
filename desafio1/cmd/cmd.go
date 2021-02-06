package cmd

import (
	"os"

	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio1/application/commands"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio1/infrastructure/db"
	"github.com/spf13/cobra"
)

var userName string
var userEmail string

var addUserCMD = &cobra.Command{
	Use:   "adduser",
	Short: "Run adduser for data generation",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		defer database.Close()

		commands.TestAddUser(database, userName, userEmail)

	},
}

func init() {
	rootCmd.AddCommand(addUserCMD)
	addUserCMD.Flags().StringVarP(&userName, "name", "n", "", "User name")
	addUserCMD.Flags().StringVarP(&userEmail, "email", "e", "", "User email")
}
