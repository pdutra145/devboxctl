package inputs

import (
	"bufio"
	"devboxctl/cli"
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

type State = int

const (
	Continue State = 0
	Exit State = 1
)

func keyboardChoiceInput(options *Choices) State {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer keyboard.Close()

	choices := *options

	char, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	if key == keyboard.KeyArrowUp && selectedIndex > 0 {
		selectedIndex--
	} else if key == keyboard.KeyArrowDown && selectedIndex < len(*options)-1 {
		selectedIndex++
	} else if key == keyboard.KeyEnter {
		choice = choices[selectedIndex]
		return Exit
	}

	if char == 'q' || char == 'Q' {
		cli.Alert.Println("Exiting...")
		return Exit
	}

	return Continue
}

func GetUserInput(message string, inputVar *string) {
	scanner := bufio.NewScanner(os.Stdin)

	cli.Cyan.Print(message)
	scanner.Scan()
	*inputVar = scanner.Text()
	
	fmt.Println()
}

func AddContainerInput() {
	var name string
	
	GetUserInput("Name of Dev Container: ", &name)

	cli.Cyan.Println("Specify the Path of .devcontainer")
	choice := DisplayChoices(Choices{0:"Current Dir", 1:"Other"})

	var path string
	switch choice {
	case 0:
		wd, err := os.Getwd()

		if err != nil {
			log.Fatal(cli.Alert.Sprint("Error in getting current dir"))
		}

		path = wd
	case 1:
		GetUserInput("Enter the path: ", &path)
	default:
		path = ""
	}

}