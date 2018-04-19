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
	"github.com/docker/docker/api/types/network"
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
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"web": &network.EndpointSettings{
				Aliases:[]string{docli.ContainerName},
				NetworkID:"98a8dacee24757b5b060ea7a03bfc9e2f00d8a3faca93b6b39f9034390eb4044",
			},
		},
	}
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, networkConfig, docli.ContainerName)
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

func ConnectToNetwork(containerId string, networkId string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("error creating docker client")
	}
	err = cli.NetworkConnect(context.Background(), networkId, containerId, &network.EndpointSettings{

	})
	if err != nil {
		fmt.Println(err)
		return errors.New("error connecting container to network")
	}
	return nil
}


func generateConfigs(docli models.DocliObject) (*container.Config, *container.HostConfig, error) {
	exposedPorts, portBindings, err := createPorts(docli)
	if err != nil {
		return &container.Config{}, &container.HostConfig{}, err
	}
	url := "Host:" + docli.ContainerName + ".valas.cloud"
	labels := map[string]string{
		"traefik.backend": docli.ContainerName,
		"traefik.docker.network": "web",
		"traefik.frontend.rule": url,
		"traefik.enable": "true",
		"traefik.port": strconv.Itoa(docli.ServerPorts[0].Host),
	}
	config := &container.Config {
		Image: docli.FullName,
		ExposedPorts: exposedPorts,
		Labels: labels,
	}
	hostConfig := &container.HostConfig {
		Binds: []string{},
		PortBindings: portBindings,
		NetworkMode: "web",
		RestartPolicy: container.RestartPolicy{
			Name: "always",
			MaximumRetryCount: 0,
		},
		VolumesFrom: []string{},

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
		//portBindings[port] = nil
		portBindings[port] = []nat.PortBinding {
			{
				HostIP: "0.0.0.0",
				HostPort: strconv.Itoa(serverPort.Host),
			},
		}
	}
	return exposedPorts, portBindings, nil
}