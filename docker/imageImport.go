package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"fmt"
	"os"
	"github.com/hengel2810/api_docli/models"
)

func ImportDockerImage(requestDockerImage models.RequestDockerImage) error {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return err
	}
	reader, err := os.Open(requestDockerImage.Path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp, err := dockerClient.ImageLoad(context.Background(), reader, false)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}