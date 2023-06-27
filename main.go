package main

import (
	"devboxctl/commands"
	"devboxctl/handler"
)

func main() {
	if err := commands.App.Execute(); err != nil {
		handler.Fatal("Unable to execute devboxctl", err)
	}
}