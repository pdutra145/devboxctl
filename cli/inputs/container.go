package inputs

import (
	"devboxctl/cli"
	"devboxctl/utils"
	"log"
	"os"
)

func AddContainerInput(info *utils.ContainerInfo) {
	name := GetUserInput("Name of Dev Container: ")
	info.Name = name

	cli.Cyan.Println("Specify the Path of .devcontainer")
	choice := DisplayChoices(Choices{0:"Current Dir", 1:"Other"})

	switch choice {
	case 0:
		wd, err := os.Getwd()

		if err != nil {
			log.Fatal(cli.Alert.Sprint("Error in getting current dir"))
		}

		info.Path = wd
	case 1:
		path := GetUserInput("Enter the path: ")
		info.Path = path
	default:
		info.Path = ""
	}

}