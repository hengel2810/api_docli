package docker

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
	"github.com/Pallinder/sillyname-go"
	"strings"
	"github.com/hengel2810/api_docli/models"
	"errors"
)

func StartContainer(docli models.DocliObject) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("error creatind docker client")
	}
	containerName := sillyname.GenerateStupidName()
	containerName = strings.ToLower(containerName)
	containerName = strings.Replace(containerName, " ", "-", -1)
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: docli.FullName,
	}, &container.HostConfig{
	}, nil, containerName)
	if err != nil {
		return errors.New("error creatind docker container")
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return errors.New("error starting docker container")
	}
	return nil
}
