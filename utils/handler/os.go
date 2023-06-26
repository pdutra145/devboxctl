package handler

import (
	"devboxctl/utils"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

type ContainerInfo struct {
	Name string
	Path string
	Image string
	DockerCompose string ""
	DockerFile string ""
}

type DevContainers = []ContainerInfo

type FilePath = string

func FileExists(file FilePath) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true // File exists
	}
	if os.IsNotExist(err) {
		return false // File does not exist
	}
	return false // Error occurred while checking file existence
}

func CreateDevContainersFile(filePath string) {
	file, err := os.Create(filePath)

	if err != nil {
		utils.Alert.Println("An error occured when trying to create file")
		return
	}

	file.WriteString("[]")

	defer file.Close()

	parts := strings.Split(filePath, "/")

	name := parts[len(parts)-1] 

	utils.Success.Printf("%s Successfully Created !\n", utils.Special.Sprint(name))
}

func CreateDir(dirPath string, mode fs.FileMode) {

	err := os.MkdirAll(dirPath, mode)

	if err != nil {
		Fatal("Unable to create Directory", err)
	}
}

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

func WriteJson(data []byte, path string) {
	file, openErr := os.Create(path)
	if openErr != nil {
		log.Fatalln("Error opening file:", openErr)
	}

	defer file.Close()

	_, fileErr := file.Write(data)
	if fileErr != nil {
		log.Fatalln("Error writing to file:", fileErr)
	}
}

