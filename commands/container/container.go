package container

import (
	"devboxctl/handler"
	"devboxctl/utils"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	name := args[0]

	fmt.Println(name)
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New(utils.Warning.Sprintf("provide the devcontainer name"))
	}

	devcontainers := handler.ReadDevContainersFile("data/devcontainers.json")

	name := args[0]

	for _, container := range devcontainers {
		if name == container.Name {
			return nil
		}
	}

	return errors.New(utils.Warning.Sprintf("invalid dev container \"%s\"\n", name))
}

var (
	Container *cobra.Command = &cobra.Command{
		Use: "container",
		Short: "Command to access properties and control your devcontainers",
		Run:run,
		Args:args,
	}
)

func init() {
	// Container.Args = cobra.ExactArgs(1)
	Container.Example = "devboxctl container test"
}