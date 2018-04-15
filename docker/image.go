package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"fmt"
	"io"
	"os"
)

func PullImage(image string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return  false
	}
	closer, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{All: true, RegistryAuth:"123"})
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = io.Copy(os.Stdout, closer)
	if err != nil {
		fmt.Println(err)
		return false
	}
	closer.Close()
	return true
}
