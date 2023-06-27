package handler

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadDevContainersFile(filePath string) DevContainers {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	// Declare a variable of the struct type
	var fileContent DevContainers

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(content, &fileContent)
	if err != nil {
		log.Fatalln("Error decoding JSON:", err)
	}

	return fileContent
}