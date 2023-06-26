package inputs

import (
	"devboxctl/utils"
	"devboxctl/utils/handler"
	"log"
	"os"
)

func AddContainerInput(info *handler.ContainerInfo) {
	name := GetUserInput("Name of Dev Container: ")
	info.Name = name

	utils.Normal.Println("Specify the Path of .devcontainer")
	choice := DisplayChoices(Choices{0:"Current Dir", 1:"Other"})

	switch choice {
	case 0:
		wd, err := os.Getwd()

		if err != nil {
			log.Fatal(utils.Alert.Sprint("Error in getting current dir"))
		}

		info.Path = wd
	case 1:
		path := GetUserInput("Enter the path: ")
		info.Path = path
	default:
		info.Path = ""
	}

	info.Image = GetUserInput("Enter the devcontainer Image: ")
}