package handler

import (
	"devboxctl/utils"
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

func AppendToJson(path FilePath, info *ContainerInfo) []ContainerInfo {
	content := ReadDevContainersFile(path)
	
	content = append(content, *info)
	return content
}