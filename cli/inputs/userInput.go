package inputs

import (
	"bufio"
	"devboxctl/cli"
	"fmt"
	"os"
)



func GetUserInput(message string, inputVar *string) {
	scanner := bufio.NewScanner(os.Stdin)

	cli.Cyan.Print(message)
	scanner.Scan()
	*inputVar = scanner.Text()

	fmt.Println()
}

