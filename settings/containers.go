package settings

import (
	"fmt"
	"os"
)

var (
	containerFile string = "devcontainers.json"
)

func containerFileExists() bool {
	filename := containerFile

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist in the current directory.\n", filename)
		return false
	} 
		fmt.Printf("File '%s' exists in the current directory.\n", filename)
	return true
}

// func GetContainers() {
// 	exists := containerFileExists()

// 	if !exists {
// 		fmt.Pr
// 	}
// }