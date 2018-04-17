package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func PullImage(image string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return  err
	}
	closer, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{All: true, RegistryAuth:"123"})
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, closer)
	if err != nil {
		return err
	}
	closer.Close()
	return nil
}
