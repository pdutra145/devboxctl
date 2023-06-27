package commands

/*
All the Commands of the Devboxctl are imported and added here
*/

import (
	add "devboxctl/commands/add"
	container "devboxctl/commands/container"

	"github.com/spf13/cobra"
)

var App *cobra.Command = &cobra.Command{
	Use:"devboxctl",
	Short: "\"devboxctl\" is a cli app that help manage devcontainers",
}

func init() {
	App.AddCommand(add.Add)
	App.AddCommand(container.Container)
}