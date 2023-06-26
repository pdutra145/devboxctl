package commands

import (
	"devboxctl/cli"
	"devboxctl/cli/inputs"
	"devboxctl/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	devContainersName string = "devcontainers.json"
	devContainersDir string = "settings/data"
	currDir string
	devContainersPath string
)

func init() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatalln(cli.Alert.Sprint("Unable to get Working Dir"))
	}

	currDir = wd

	devContainersPath = filepath.Join(currDir, devContainersDir, devContainersName)
}



func AddContainer() utils.ContainerInfo {
	boldGreen := cli.Confirm.Add(color.Bold)

	if !(utils.FileExists(devContainersPath)) {
		utils.CreateFile(devContainersPath, devContainersName)
	} 
	var info utils.ContainerInfo
	inputs.AddContainerInput(&info)

	content := utils.ReadJsonFile(devContainersPath)
	
	content = append(content, info)
	fmt.Println(content)

	contentJson, err := json.Marshal(content)

	if err != nil {
		log.Fatalln(cli.Alert.Sprint("Error in converting to Json"))
	}

	utils.WriteJson(contentJson, devContainersPath)

	boldGreen.Println("Content added to devcontainers.json")

	return info
}

func AddCreateContainer() {
	info := AddContainer()

	filePath := filepath.Join(info.Path, ".devcontainer")

	file, err := os.Create(filePath)

	if err != nil {
		log.Fatalln(cli.Alert.Sprint("Error in creating the .devcontainers folder"))
	}

	defer file.Close()


}