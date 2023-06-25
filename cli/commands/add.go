package commands

import (
	"devboxctl/cli"
	"fmt"
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

	defer file.Close()

	cli.Confirm.Printf("%s Successfully Created !\n", cli.Cyan.Sprint(fileName))
}

func AddContainer() {
	_, err := os.Stat(filePath)
	fmt.Print(err)

	if os.IsNotExist(err) {
		createFile()
	}
}