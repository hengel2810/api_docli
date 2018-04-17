package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
	"os"
	"errors"
)

func PullImage(image string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("error creating docker client")
	}
	closer, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{All: true, RegistryAuth:"123"})
	if err != nil {
		return errors.New("error pulling docker image")
	}
	_, err = io.Copy(os.Stdout, closer)
	if err != nil {
		return errors.New("error copying docker image")
	}
	closer.Close()
	return nil
}

func RemoveImage(image string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("error creating docker client")
	}
	_, err = cli.ImageRemove(context.Background(), image, types.ImageRemoveOptions{Force:true, PruneChildren:false})
	if err != nil {
		return errors.New("error removing docker image")
	}
	return nil
}
