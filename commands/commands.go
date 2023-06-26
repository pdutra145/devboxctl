package commands

import (
	"github.com/spf13/cobra"
)

var App *cobra.Command = &cobra.Command{
	Use:"devboxctl",
	Short: "\"devboxctl\" is a cli app that help manage devcontainers",
}

func init() {
	App.AddCommand(Add)
}