package container

import (
	"devboxctl/handler"
	"devboxctl/utils"
	"fmt"
	"os/exec"
)

func CheckIfContainerExists(name string) bool {
	devcontainers := handler.ReadDevContainersFile()

	var containerExists bool
	for _, container := range devcontainers {
		if name == container.Name {
			containerExists = true
		}
	}

	return containerExists
}

func GetContainerInfo(name string, info *handler.ContainerInfo) {
	devcontainers := handler.ReadDevContainersFile()

	for _, container := range devcontainers {
		if name == container.Name {
			*info = container
			return
		}
	}
}

func DevContainerUp(info *handler.ContainerInfo) {
	cmd := exec.Command("docker-compose", "-f", info.Path, info.Name, "up", "-d")

	res, err := cmd.Output()

	if err != nil {
		handler.Fatal(utils.Alert.Sprint("Unable to start container"), err)
	}

	fmt.Println(res)
}