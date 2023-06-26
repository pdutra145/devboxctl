package inputs

import (
	"bufio"
	"devboxctl/utils"
	"fmt"
	"os"
)

type Message = string

func GetUserInput(message Message) string {
	var response string
	scanner := bufio.NewScanner(os.Stdin)

	utils.Special.Print(message)
	scanner.Scan()
	response = scanner.Text()

	fmt.Println()
	return response
}

func GetBooleanUserInput(message Message) bool {
	options := Choices{
		0: "No",
		1: "Yes",
	}
	
	result := DisplayChoices(options)

	return options[result] == "Yes"
}