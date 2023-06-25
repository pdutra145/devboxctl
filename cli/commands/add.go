package commands

import (
	"devboxctl/cli"
	"devboxctl/cli/inputs"
	"devboxctl/utils"
	"os"
	"path/filepath"
)

var (
	fileName string = "devcontainers.json"
	fileDir string = "settings/data"
	filePath string = filepath.Join(fileDir, fileName)
)

func createFile() {
	file, err := os.Create(filePath)

	if err != nil {
		cli.Alert.Println("An error occured when trying to create file")
		return
	}

	file.WriteString("{}")

	defer file.Close()

	cli.Confirm.Printf("%s Successfully Created !\n", cli.Cyan.Sprint(fileName))
}

func AddContainer() {
	if !(utils.FileExists(filePath)) {
		createFile()
	} 

	inputs.AddContainerInput()
}