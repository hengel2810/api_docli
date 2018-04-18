package docker

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
	"github.com/hengel2810/api_docli/models"
	"errors"
	"fmt"
	"strconv"
	"github.com/docker/go-connections/nat"
)

func StartContainer(docli models.DocliObject) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("error creating docker client")
	}
	config, hostConfig, err := generateConfigs(docli)
	if err != nil {
		return err
	}
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, docli.ContainerName)
	if err != nil {
		fmt.Println(err)
		return errors.New("error creating docker container")
	}
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		StopContainer(docli)
		return errors.New("error starting docker container")
	}
	return nil
}

func StopContainer(docli models.DocliObject) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("error creating docker client")
	}
	err = cli.ContainerRemove(context.Background(), docli.ContainerName, types.ContainerRemoveOptions{
		RemoveLinks:false,
		RemoveVolumes:false,
		Force:true,
	})
	if err != nil {
		return errors.New("error removing docker container")
	}
	return nil
}


func generateConfigs(docli models.DocliObject) (*container.Config, *container.HostConfig, error) {
	exposedPorts, portBindings, err := createPorts(docli)
	if err != nil {
		return &container.Config{}, &container.HostConfig{}, err
	}
	config := &container.Config {
		Image: docli.FullName,
		ExposedPorts: exposedPorts,
	}

	hostConfig := &container.HostConfig {
		//Binds: []string{
		//	"/var/run/docker.sock:/var/run/docker.sock",
		//},
		PortBindings: portBindings,
	}
	return config, hostConfig, nil
}

func createPorts(docli models.DocliObject) (nat.PortSet, nat.PortMap, error) {
	exposedPorts := nat.PortSet{}
	portBindings := nat.PortMap{}
	for _, serverPort := range docli.ServerPorts {
		port, err := nat.NewPort("tcp", strconv.Itoa(serverPort.Container))
		if err != nil {
			return exposedPorts, portBindings, errors.New("error creating port")
		}
		exposedPorts[port] = struct{}{}
		portBindings[port] = []nat.PortBinding {
			{
				HostIP: "0.0.0.0",
				HostPort: strconv.Itoa(serverPort.Host),
			},
		}
	}
	return exposedPorts, portBindings, nil
}