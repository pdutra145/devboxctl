package inputs

import (
	"bufio"
	"devboxctl/cli"
	"fmt"
	"os"
)

type UserResponse = string

func GetUserInput(message string) UserResponse {
	var response UserResponse
	scanner := bufio.NewScanner(os.Stdin)

	cli.Cyan.Print(message)
	scanner.Scan()
	response = scanner.Text()

	fmt.Println()
	return response
}

