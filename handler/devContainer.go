package handler

import (
	"devboxctl/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type DevContainerJSON interface {
	Marshal() []byte
}

type DevContainer struct {
	Name string `json:"name"`
	DockerComposeFile string `json:"dockerComposeFile"`
	Service string `json:"service"`
	WorkspaceFolder string `json:"workspaceFolder"`
	ShutdownAction string  `json:"shutdownAction"`
}

func (c *DevContainer) Marshal() []byte {
	data, err := json.Marshal(c)

	if err != nil {
		Fatal("Error in encoding devcontainer.json data", err)
	}

	return data
}


func CreateDevContainerFile(filePath FilePath, info *ContainerInfo) {
	var jsonInfo DevContainerJSON = &DevContainer{
		Name: info.Name,
		DockerComposeFile: info.DockerCompose,
		Service: "devcontainer",
		WorkspaceFolder: "/workspace",
		ShutdownAction: "stopCompose",
	}

	data := jsonInfo.Marshal()

	file, err := os.Create(filePath)

	if err != nil {
		Fatal("Unable to create devcontainer.json", err)
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		Fatal("Unable to write to devcontainer.json", err)
	}
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

func CreateDockerFile(filePath FilePath, info ContainerInfo) {
	file, err := os.Create(filePath)

	if err != nil {
		Fatal("Unable to create Dockerfile", err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("FROM %s", info.Image))

	if err != nil {
		Fatal("Unable to write to Dockerfile", err)
	}

	utils.Success.Println("Dockerfile Successfully Created")
}

func CreateEnvFile(filePath FilePath, info ContainerInfo) {
	file, err := os.Create(filePath)

	if err != nil {
		Fatal("Unable to Create Env file", err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("CONTAINER_NAME=%s", info.Name))
	
	if err != nil {
		Fatal("Unable to Write to Env File", err)
	}

	utils.Success.Println("Env file Successfully Created!")
}