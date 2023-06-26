package commands

import (
	"devboxctl/inputs"
	"devboxctl/utils"
	"devboxctl/utils/handler"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("new") {
		AddCreateContainer()
	}
	AddContainer()
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
	devContainersDir := "settings/data"

	devContainersPath = filepath.Join(currDir, devContainersDir, devContainersName)

	Add.Flags().StringP("new", "n", "", "add to devcontainers.json and setup the devcontainer")
}

func AddContainer() handler.ContainerInfo {
	boldGreen := utils.Success.Add(color.Bold)

	if !(handler.FileExists(devContainersPath)) {
		handler.CreateFile(devContainersPath)
	} 
	var info handler.ContainerInfo
	inputs.AddContainerInput(&info)

	content := handler.ReadJsonFile(devContainersPath)
	
	content = append(content, info)
	fmt.Println(content)

	contentJson, err := json.Marshal(content)

	if err != nil {
		handler.Fatal("Error in converting to Json", err)
	}

	handler.WriteJson(contentJson, devContainersPath)

	boldGreen.Println("Content added to devcontainers.json")

	return info
}

func AddCreateContainer() {
	info := AddContainer()

	dirPath := filepath.Join(info.Path, ".devcontainer")

	handler.CreateDir(dirPath, os.ModePerm)

	createCompose := inputs.GetBooleanUserInput("Do you want to configure docker-compose.yml file ? ")

	if createCompose {
		filePath := filepath.Join(dirPath, "docker-compose.yml")
		handler.CreateFile(filePath)
	}

	createDockerFile := inputs.GetBooleanUserInput("Do you want to configure a Dockerfile ? ")

	if createDockerFile {
		filePath := filepath.Join(dirPath, "Dockerfile")
		handler.CreateFile(filePath)
	}
}

