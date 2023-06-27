package commands

import (
	"devboxctl/inputs"
	"devboxctl/utils"
	"devboxctl/utils/handler"
	"encoding/json"
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

func AddContainer() handler.ContainerInfo {
	if !(handler.FileExists(devContainersPath)) {
		handler.CreateDevContainersFile(devContainersPath)
	} 
	var info handler.ContainerInfo
	inputs.AddContainerInput(&info)

	content := handler.AppendToJson(devContainersPath, &info)

	contentJson, err := json.Marshal(content)

	if err != nil {
		handler.Fatal("Error in converting to Json", err)
	}

	handler.WriteJson(contentJson, devContainersPath)

	utils.Success.Println("Content added to devcontainers.json")

	return info
}

func AddCreateContainer() {
	info := AddContainer()

	dirPath := filepath.Join(info.Path, ".devcontainer")

	handler.CreateDir(dirPath, os.ModePerm)

	var filePath string

	filePath = filepath.Join(dirPath, "docker-compose.yml")
	handler.CreateDockerComposeFile(filePath)
	

	filePath = filepath.Join(dirPath, "Dockerfile")
	handler.CreateDockerFile(filePath, info)

	filePath = filepath.Join(dirPath, "devcontainer.json")
	handler.CreateDevContainerFile(filePath, info)

	filePath = filepath.Join(dirPath, ".env")
	handler.CreateEnvFile(filePath, info)

	utils.Special.Printf("\nDev container ready to use in %s\n", utils.Normal.Sprint(dirPath))
}

