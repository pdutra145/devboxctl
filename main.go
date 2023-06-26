package main

import (
	"devboxctl/commands"
	"devboxctl/utils/handler"
)

func main() {
	if err := commands.App.Execute(); err != nil {
		handler.Fatal("Unable to execute devboxctl", err)
	}
}