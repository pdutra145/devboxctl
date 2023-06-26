package handler

import (
	"devboxctl/utils"
	"encoding/json"
	"io"
	"os"
)

type DevContainerJSON interface {
	Marshal() []byte
}

type DevContainer struct {
	name string
	dockerComposeFile string
	service string
	workspaceFolder string
	shutdownAction string 
}

func (c *DevContainer) Marshal() []byte {
	data, err := json.Marshal(c)

	if err != nil {
		Fatal("Error in encoding devcontainer.json data", err)
	}

	return data
}


func CreateDevContainerFile(filePath FilePath, info ContainerInfo) {
	var jsonInfo DevContainerJSON = &DevContainer{
		name: info.Name,
		dockerComposeFile: info.DockerCompose,
		service: "devcontainer",
		workspaceFolder: "/workspace",
		shutdownAction: "stopCompose",
	}

	data := jsonInfo.Marshal()

	file, err := os.Create(filePath)

	if err != nil {
		Fatal("Unable to create devcontainer.json", err)
	}

	defer file.Close()

	WriteJson(data, filePath)
}

func CreateDockerComposeFile(filePath FilePath) {
	templatePath := "data/docker-compose.template.yml"

	dockerComposeFile, err := os.Create(filePath)

	if err != nil {
		Fatal("Unable to create docker-compose.yml", err)
	}

	defer dockerComposeFile.Close()

	templateFile, err := os.Open(templatePath)

	if err != nil {
		Fatal("Unable to read docker-compose.yml template", err)
	}

	defer templateFile.Close()

	_, err = io.Copy(dockerComposeFile, templateFile)
	if err != nil {
		Fatal("Unable to copy contents from template to docker-compose.yml", err)
	}

	utils.Success.Println("Docker Compose file Created Successfully")
}