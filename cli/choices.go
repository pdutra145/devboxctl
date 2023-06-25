package cli

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var (
	redBold *color.Color = color.New(color.FgRed, color.Bold)
	cyan *color.Color = color.New(color.FgCyan)
	yellow *color.Color = color.New(color.FgYellow)
	green *color.Color = color.New(color.FgGreen)
	choice string
)

type Choices = []string

func printOptions(options Choices, selectedIndex int) {
	boldYellow := yellow.Add(color.Bold).Sprint


	cyan.Printf("Use the arrow keys to select an option %s", boldYellow("(Press Enter to confirm):\n"))

	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("%s %s\n", cyan.Sprint(">"), option)
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

	selectedIndex := 0
	green := green.Sprint

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer keyboard.Close()

	for {
		printOptions(options, selectedIndex)

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowUp && selectedIndex > 0 {
			selectedIndex--
		} else if key == keyboard.KeyArrowDown && selectedIndex < len(options)-1 {
			selectedIndex++
		} else if key == keyboard.KeyEnter {
			choice = options[selectedIndex]
			break
		}

		if char == 'q' || char == 'Q' {
			redBold.Println("Exiting...")
			return
		}
		clearScreen()
	}


	fmt.Printf("\n%s %s\n", green("Selected Option: "), choice)
}