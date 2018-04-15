package docker

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
	"github.com/hengel2810/api_docli/models"
)

func StartContainer(requestDockerImage models.RequestDockerImage) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return  err
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: requestDockerImage.FullName,
	}, nil, nil, "abc1")
	if err != nil {
		return  err
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return  err
	}
	return nil
}
