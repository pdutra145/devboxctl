package main

import (
	"devbox/cli"
)


func main() {
	choices := []string{"Option 1", "Option 2"}

	cli.DisplayChoices(choices)
}