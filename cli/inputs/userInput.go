package inputs

import (
	"devboxctl/cli"
	"log"

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