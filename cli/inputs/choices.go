package inputs

import (
	"devboxctl/cli"
	"devboxctl/utils"
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

type Choice string
type Choices = map[int]Choice

var (
	selectedIndex int = 0
)


func printOptions(options Choices, selectedIndex int) {
	boldYellow := cli.Yellow.Add(color.Bold).Sprint
	cyan := cli.Cyan


	cyan.Printf("Use the arrow keys to select an option %s", boldYellow("(Press Enter to confirm):\n"))

	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("%s %s\n", cyan.Sprint(">"), option)
		} else {
			fmt.Println(option)
		}
	}
}

func DisplayChoices(options Choices) int {
	if !(len(options) > 0) {
		log.Fatal("Options is Empty")
	}

	for {
		printOptions(options, selectedIndex)

		state := keyboardChoiceInput(&options)

		if state == 1 {
			break
		}

		utils.ClearScreen()
	}


	return selectedIndex
}

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


	char, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	if key == keyboard.KeyArrowUp && selectedIndex > 0 {
		selectedIndex--
	} else if key == keyboard.KeyArrowDown && selectedIndex < len(*options)-1 {
		selectedIndex++
	} else if key == keyboard.KeyEnter {
		return Exit
	}

	if char == 'q' || char == 'Q' {
		cli.Alert.Println("Exiting...")
		return Exit
	}

	return Continue
}

