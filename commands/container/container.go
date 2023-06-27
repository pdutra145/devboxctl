package container

import (
	"devboxctl/handler"
	"devboxctl/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	name := args[0]
	action := strings.ToLower(args[1])

	var container handler.ContainerInfo
	GetContainerInfo(name, &container)

	if action == "up" {
		DevContainerUp(&container)
	}

	fmt.Println(name)
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New(utils.Warning.Sprintf("Insufficient arguments"))
	}


	name := args[0]

	containerExists := CheckIfContainerExists(name)

	if !containerExists {
		return errors.New(utils.Warning.Sprintf("invalid dev container \"%s\"\n", name))
	}

	return nil
}

var (
	Container *cobra.Command = &cobra.Command{
		Use: "container",
		Short: "Command to access properties and control your devcontainers",
		Run:run,
		Args:args,
	}
	usageDoc string = `Usage: 
	devboxctl container <containerName> [flags]
  
  Flags:
	-h, --help for help
	  `
)

func init() {
	// Container.Args = cobra.ExactArgs(1)
	Container.Example = "devboxctl container test"
	Container.SetUsageTemplate(usageDoc)
}