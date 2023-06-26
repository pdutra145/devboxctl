package utils

import (
	"devboxctl/cli"
	"encoding/json"
	"io"
	"log"
	"os"
)

type ContainerInfo struct {
	Name string
	Path string
	DockerCompose string ""
	DockerFile string ""
}

type DevContainers = []ContainerInfo

func CreateFile(path string, name string) {
	file, err := os.Create(path)

	if err != nil {
		cli.Alert.Println("An error occured when trying to create file")
		return
	}

	file.WriteString("[]")

	defer file.Close()

	cli.Confirm.Printf("%s Successfully Created !\n", cli.Cyan.Sprint(name))
}

func ReadJsonFile(filePath string) DevContainers {
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
			log.Fatalln("Error unmarshaling JSON:", err)
		}
	
		return fileContent
}

func WriteJson(data []byte, path string) {
	file, error1 := os.Create(path)
	if error1 != nil {
		log.Fatalln("Error opening file:", error1)
	}

	defer file.Close()

	_, err := file.Write(data)
	if err != nil {
		log.Fatalln("Error writing to file:", err)
	}
}