package add

import (
	"devboxctl/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("new") {
		AddCreateContainer()
	} else {
		AddContainer()
	}
}

var (
	devContainersPath string
	Add *cobra.Command = &cobra.Command{
		Use: "add",
		Short:"\"add\" adds containers to devcontainers.json",
		Run : run,
	}
)

func init() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatalln(utils.Alert.Sprint("Unable to get Working Dir"))
	}

	currDir := wd
	devContainersName := "devcontainers.json"
	devContainersDir := "data"

	devContainersPath = filepath.Join(currDir, devContainersDir, devContainersName)

	Add.Flags().BoolP("new", "n", false, "add to devcontainers.json and setup the devcontainer")
}



