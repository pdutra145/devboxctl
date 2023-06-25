package inputs

import (
	"devboxctl/cli"
	"log"
	"os"
)

func AddContainerInput(name *string, path *string) {
	
	GetUserInput("Name of Dev Container: ", name)

	cli.Cyan.Println("Specify the Path of .devcontainer")
	choice := DisplayChoices(Choices{0:"Current Dir", 1:"Other"})

	switch choice {
	case 0:
		wd, err := os.Getwd()

		if err != nil {
			log.Fatal(cli.Alert.Sprint("Error in getting current dir"))
		}

		*path = wd
	case 1:
		GetUserInput("Enter the path: ", path)
	default:
		*path = ""
	}

}