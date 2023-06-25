package inputs

import (
	"devboxctl/cli"
	"fmt"
	"log"

	"github.com/fatih/color"
)

var (
	choice string
	selectedIndex int = 0
)

type Choices = []string

func printOptions(options Choices, selectedIndex int) {
	boldYellow := cli.Yellow.Add(color.Bold).Sprint


	cli.Cyan.Printf("Use the arrow keys to select an option %s", boldYellow("(Press Enter to confirm):\n"))

	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("%s %s\n", cli.Cyan.Sprint(">"), option)
		} else {
			fmt.Println(option)
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape sequence to clear screen
}

func DisplayChoices(options Choices) {
	if !(len(options) > 0) {
		log.Fatal("Options is Empty")
	}

	
	Green := cli.Confirm.Sprint

	

	for {
		printOptions(options, selectedIndex)

		state := keyboardChoiceInput(&options)

		if state == 1 {
			break
		}

		clearScreen()
	}


	fmt.Printf("\n%s %s\n", Green("Selected Option: "), choice)
}