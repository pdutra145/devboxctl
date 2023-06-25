package main

import (
	"devboxctl/cli"
	"devboxctl/cli/commands"
	"flag"

	"log"
)


func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal(cli.Alert.Sprint("No Command Provided"))
	}

	// choices := []string{"Option 1", "Option 2"}

	command := args[0]

	switch command {
	case "add":
		commands.AddContainer()
	}

	// cli.DisplayChoices(choices)
}