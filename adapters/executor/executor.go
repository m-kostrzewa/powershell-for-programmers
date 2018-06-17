package executor

import (
	"context"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"github.com/satori/go.uuid"
)

type PowershellDockerExecutor struct {
	dockerClient *docker.Client
}

func New() (*PowershellDockerExecutor, error) {
	pde := &PowershellDockerExecutor{}
	// client, err := docker.NewClient("http://127.0.0.1", "1", nil, nil)
	client, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}
	pde.dockerClient = client
	return pde, nil
}

func (executor *PowershellDockerExecutor) Execute(code string) (string, error) {
	cfg := &container.Config{
		Image:        "mcr.microsoft.com/powershell",
		Entrypoint:   []string{"powershell.exe", "Write-Host", "'hi'"},
		AttachStdout: true,
	}
	name := uuid.NewV4().String()[:5]
	resp, err := executor.dockerClient.ContainerCreate(context.Background(),
		cfg, nil, nil, name)
	if err != nil {
		return "", err
	}
	containerID := resp.ID
	defer executor.dockerClient.ContainerRemove(context.Background(),
		containerID, types.ContainerRemoveOptions{Force: true})

	err = executor.dockerClient.ContainerStart(context.Background(),
		containerID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	executor.dockerClient.ContainerLogs

	return "", err
}
