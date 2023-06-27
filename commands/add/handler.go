package add

import (
	"devboxctl/handler"
	"devboxctl/inputs"
	"devboxctl/utils"
	"encoding/json"
	"os"
	"path/filepath"
)

func AddContainer() handler.ContainerInfo {
	if !(handler.FileExists(devContainersPath)) {
		handler.CreateDevContainersFile(devContainersPath)
	} 
	
	var info handler.ContainerInfo
	inputs.AddContainerInput(&info)

	content := handler.AppendToJson(devContainersPath, &info)

	contentJson, err := json.Marshal(content)

	if err != nil {
		handler.Fatal("Error in converting to Json", err)
	}

	handler.WriteJson(contentJson, devContainersPath)

	utils.Success.Println("Content added to devcontainers.json")

	return info
}

func AddCreateContainer() {
	info := AddContainer()

	dirPath := filepath.Join(info.Path, ".devcontainer")

	handler.CreateDir(dirPath, os.ModePerm)

	var filePath string

	filePath = filepath.Join(dirPath, "docker-compose.yml")
	handler.CreateDockerComposeFile(filePath)
	

	filePath = filepath.Join(dirPath, "Dockerfile")
	handler.CreateDockerFile(filePath, info)

	filePath = filepath.Join(dirPath, "devcontainer.json")
	handler.CreateDevContainerFile(filePath, &info)

	filePath = filepath.Join(dirPath, ".env")
	handler.CreateEnvFile(filePath, info)

	utils.Special.Printf("\nDev container ready to use in %s\n", utils.Normal.Sprint(dirPath))
}
